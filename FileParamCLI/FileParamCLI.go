package FileParamCLI

import (
	"bufio"
	"fmt"
	"os"
	"rabadaz.at/defaultfiles/Config"
	"strings"
)

type ReplacementParameters struct{
	Parameter Config.FileParamenter
	Replacement string
}

func DecideParams(params []Config.FileParamenter) []ReplacementParameters{
	paramLen := len(params)
	RpParams := make([]ReplacementParameters, paramLen)

	reader := bufio.NewReader(os.Stdin)

	for i, p := range params{
		fmt.Println("\n=========\nParameter to set:", p.Description)
		fmt.Printf("Set value of Parameter [%s]: ", p.Key)
		text, _ := reader.ReadString('\n')

		text = strings.ReplaceAll(text, "\n", "")

		RpParams[i] = ReplacementParameters{
			Replacement: text,
			Parameter: p,
		}
	}
	return RpParams
}