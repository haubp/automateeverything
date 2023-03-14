package nodeBiz

import (
	"backend_autotest/common"
	"os"
)

func DeleteAutoLogFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {

		return common.ErrCannotDeleteEntity("fileLogAuto", err)
	}
	return nil
}
