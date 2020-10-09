package processors

import (
	"fmt"

	"github.com/example/app/common"
)

func Process(line string, r *common.Record) {
	r.Log += fmt.Sprintf("processed by proc1.Process(%s)\n", line)
}
