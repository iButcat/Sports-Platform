package main


import (
   "fmt"
   "os"
   "bufio"
   "github.com/jinzhu/gorm"
   _"github.com/jinzhu/gorm/dialects/postgres"
   _ "github.com/lib/pq"
)

type DbController struct {
  db *gorm.DB
}

type DbIdentifier struct {
   host, port, user, dbname, password string
}

func printFields(message string) {
  fmt.Println(message)
}

func (dbIdentifier *DbIdentifier) readStringInput() {
  reader := bufio.NewReader(os.Stdin)
  printFields("Host ->")
  dbIdentifier.host, _ = reader.ReadString('\n')
  printFields("Port ->")
  dbIdentifier.port, _ = reader.ReadString('\n')
  printFields("User ->")
  dbIdentifier.user, _ = reader.ReadString('\n')
  printFields("Db Name->")
  dbIdentifier.dbname, _ = reader.ReadString('\n')
  printFields("Password ->")
  dbIdentifier.password, _ = reader.ReadString('\n')
}

func (dc *DbController) connectDatabase() {
  var err error
  dbIdentifier := DbIdentifier{}

  dbIdentifier.readStringInput()

  connectionString := fmt.Sprintf(
    "host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
  dbIdentifier.host,
  dbIdentifier.port,
  dbIdentifier.user,
  dbIdentifier.dbname,
  dbIdentifier.password)

  dc.db, err = gorm.Open("postgres", connectionString)
  if err != nil {
      panic(err)
   } else {
     fmt.Println("Success")
   }
  dc.db.LogMode(true)
}

func (dc *DbController) retrieveDataAndSites() {
  var data Data
  rows, err := dc.db.Preload("Sites").Find(&data).Rows()
  defer rows.Close()
  if err != nil {
      panic(err)
   }
   for rows.Next() {
      dc.db.ScanRows(rows, &data)
      fmt.Println(data)
   }
}

func (dc *DbController) createModel(sports *Sports) {
   dc.db.Debug().Create(sports)
}

func (dc *DbController) dropAndAutoMigrate(sports interface{}) {
  dc.db.Debug().DropTableIfExists(sports)
  dc.db.Debug().AutoMigrate(sports)
}

func (dc *DbController) updateOdds() {
  sites := Sites{}
  _ = dc.db.Model(&sites).Select("odds").Update()
  fmt.Println("Successfully updated fields odds")
}


func (dc *DbController) deleteOldData() {
  dc.db.Exec("DELETE FROM sports, data, sites WHERE created_at < GETDATE()- 30")
}

func (dc *DbController) closeConnection() {
  defer dc.db.Close()
}
