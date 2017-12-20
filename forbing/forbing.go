package main 


import (
	"log"
	"os"
	"bufio"
	"encoding/csv"
	"math/rand"
	"time"
)
func main() {


	fk, err := os.Open("500keywords.txt")
	defer fk.Close()
	if err != nil {
		log.Fatal(err)
	}


	

	bestkeywords :=make(map[int]string)

	var kwcount int 

	scanner := bufio.NewScanner(fk)

	for scanner.Scan() {


		fline := scanner.Text()

		bestkeywords[kwcount] = fline 
		kwcount ++


	}
		
	
	ftmp, err := os.OpenFile("tmp", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644) 
	
		
	if err != nil {
		log.Fatal(err)

	}
	
	
	defer ftmp.Close()


	f, err := os.Open("domains.csv")
	
		if err != nil {
			log.Fatal(err)
	
		}

		domains := csv.NewReader(f)
		domains.FieldsPerRecord = 0
		records, err := domains.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		 	
		for _, record := range records {

			for i:=0; i <3 ; i++ {
			rand.Seed(time.Now().UTC().UnixNano())
			rnd := rand.Intn(500)

			log.Println("http://"+bestkeywords[rnd]+"."+record[2])
			_, err := ftmp.WriteString("http://"+bestkeywords[rnd]+"."+record[2]+"\n")

			if err != nil {
				log.Fatal(err)
		
			}

			}

			


		}

}