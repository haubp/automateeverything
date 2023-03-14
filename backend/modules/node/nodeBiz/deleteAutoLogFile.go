package nodeBiz

import (
	"backend_autotest/common"
	"os"
)

func DeleteAgentLogFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {

		return common.ErrCannotDeleteEntity("fileLogAgent", err)
	}
	return nil
}
