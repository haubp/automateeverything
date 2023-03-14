package template

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"net/http"
	"crypto/tls"
	"bytes"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"

	"automateeverything.com/v2/custom"
	"automateeverything.com/v2/maserver"
)

// RunTestPage Creat Run Test Page widget
func RunTestPage(a fyne.App, w fyne.Window) *fyne.Container {
	var t TestCategory
	resultTelevision := custom.NewFixSizeLabel()
	testRunProgress := widget.NewProgressBar()
	testRunProgress.SetValue(0.0)
	reportNameEntry := widget.NewEntry()
	reportNameEntry.SetText("report.json")
	templateFileSelectedLabelIndicator := widget.NewLabel("No test template found")
	scrollResultContainer := container.NewVScroll(
		resultTelevision,
	)
	runTestButton := widget.NewButton("Run", func() {
		if len(t.TestCategoryGroups) != 0 {
			ExecuteTestFromTemplate(&t, resultTelevision, testRunProgress, scrollResultContainer, reportNameEntry)
		} else {
			resultTelevision.WriteAndExpand("No Test Template is imported")
		}
	})
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	startMAServerButton := widget.NewButton("Start MA", func() {
		// TODO: Need a way to stop MA Server
		resultTelevision.WriteAndExpand("Start MA Server...")
		go func() {
			engine := maserver.InitMaServer()

			tlsConfig := &tls.Config{
				MaxVersion: tls.VersionTLS13,
				MinVersion: tls.VersionTLS13,
			}

			server := &http.Server{Addr: ":4433", Handler: engine.Handler(), TLSConfig: tlsConfig}
			err := server.ListenAndServeTLS(filepath.Join(pwd, "server.crt"), filepath.Join(pwd, "server.key"))
			if err != nil {
				log.Println(err)
			}
		}()
	})
	clearLogButton := widget.NewButton("Clear", func() {
		resultTelevision.SetText("")
	})
	importTestButton := widget.NewButton("Import", func() {
		fileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				// read files
				if r == nil {
					return
				}
				data, _ := ioutil.ReadAll(r)
				t, _ = CreateTestFromBytes(data)
				templateFileSelectedLabelIndicator.SetText("Template imported")
			}, w)
		fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))
		currentLocationURI := storage.NewFileURI(".")
		currentLocationListableURI, err := storage.ListerForURI(currentLocationURI)
		if err != nil {
			log.Println(err)
		}
		fileDialog.SetLocation(currentLocationListableURI)
		fileDialog.Show()
	})
	runTestPage := container.New(layout.NewGridLayoutWithRows(2),
		container.New(layout.NewHBoxLayout(),
			container.New(layout.NewVBoxLayout(),
				importTestButton,
				runTestButton,
				clearLogButton,
				startMAServerButton,
				layout.NewSpacer()),
			container.New(layout.NewVBoxLayout(),
				templateFileSelectedLabelIndicator,
				testRunProgress,
				reportNameEntry,
				layout.NewSpacer(),
			),
			layout.NewSpacer()),
		scrollResultContainer,
	)
	return runTestPage
}

func makeRequest(httpMethod string, body []byte, url string) []byte {
	// Send GET Command request
	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest(httpMethod, url, bodyReader)
	if err != nil {
		log.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}	
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	return resBody
}

// RunTestAutoPage Creat Run Test Page widget
func RunTestAutoPage(a fyne.App, w fyne.Window) *fyne.Container {
	var t TestCategory
	resultTelevision := custom.NewFixSizeLabel()
	testRunProgress := widget.NewProgressBar()
	testRunProgress.SetValue(0.0)
	reportNameEntry := widget.NewEntry()
	reportNameEntry.SetText("report.json")
	templateFileSelectedLabelIndicator := widget.NewLabel("No test template found")
	scrollResultContainer := container.NewVScroll(
		resultTelevision,
	)
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	resultTelevision.WriteAndExpand("Ready for downloading Test Template")

	startMAServerButton := func() {
		// TODO: Need a way to stop MA Server
		resultTelevision.WriteAndExpand("Start MA Server...")
		go func() {
			engine := maserver.InitMaServer()
			err := engine.RunTLS(":4433", filepath.Join(pwd, "server.crt"), filepath.Join(pwd, "server.key"))
			if err != nil {
				log.Println(err)
			}
		}()
	}

	autoBackendUrl := "http://10.40.50.34:6868"

	runTestButton := widget.NewButton("Start", func() {
		for {
			// Send GET Command request
			resultTelevision.WriteAndExpand("Get Node Command")
			commandBytes := makeRequest(	http.MethodGet, []byte(`{"node_id":"ec5bf52b73d078fa0ac7d27281da2104a588440e8571c2c4f00aa69d5464e387"}`), 
											autoBackendUrl + "/api/v1/node/command")
			
			// if command was RUN -> fetch test template and run
			if string(commandBytes) == "\"RUN\"" {
				resultTelevision.WriteAndExpand("RUN")

				// Get Test Template
				downloadedTemplate := makeRequest(	http.MethodGet, 
													nil, 
													autoBackendUrl + "/public/user_d82494f05d6917ba02f7aaa29689ccb444bb73f20380876cb05d1f37537b7892/template.json")

				resultTelevision.WriteAndExpand("Downloaded Template")

				t, _ = CreateTestFromBytes(downloadedTemplate)

				startMAServerButton()

				if len(t.TestCategoryGroups) != 0 {
					resultTelevision.WriteAndExpand("Downloaded Test Template")
					ExecuteTestFromTemplate(&t, resultTelevision, testRunProgress, scrollResultContainer, reportNameEntry)

					// Upload Test Result
					reportFile, err := os.Open(filepath.Join(pwd, "report.json"))
					if err != nil {
						log.Println(err)
						continue
					}
					defer reportFile.Close()
				
					reportJsonBytes, _ := ioutil.ReadAll(reportFile)
					makeRequest(	http.MethodPost, 
									reportJsonBytes, 
									autoBackendUrl + "/api/v1/node/result?node_id=ec5bf52b73d078fa0ac7d27281da2104a588440e8571c2c4f00aa69d5464e387")
				} else {
					resultTelevision.WriteAndExpand("No Test Template is imported")
				}
			} else {
				resultTelevision.WriteAndExpand("Receive " + string(commandBytes))
			}

			resultTelevision.WriteAndExpand("Ready for downloading Test Template")

			time.Sleep(time.Duration(60) * time.Second)
		}
	})
	
	clearLogButton := widget.NewButton("Clear", func() {
		resultTelevision.SetText("")
	})
	
	runTestPage := container.New(layout.NewGridLayoutWithRows(2),
		container.New(layout.NewHBoxLayout(),
			container.New(layout.NewVBoxLayout(),
				runTestButton,
				clearLogButton,
				layout.NewSpacer()),
			container.New(layout.NewVBoxLayout(),
				templateFileSelectedLabelIndicator,
				testRunProgress,
				reportNameEntry,
				layout.NewSpacer(),
			),
			layout.NewSpacer()),
		scrollResultContainer,
	)
	return runTestPage
}
