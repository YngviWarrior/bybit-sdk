package bybitSDK

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

func Log(msg any) {
	logEnv, _ := strconv.ParseBool(os.Getenv("LOGS"))

	out, err := json.MarshalIndent(msg, "", "  ") // pretty print
	if err != nil {
		log.Println("Erro ao serializar log:", err)
		return
	}

	log.Println(string(out))
	if logEnv {
		log.Println(msg)
	}
}
