package main

import (
	"flag"
	_ "net/http/pprof"

	//"github.com/golang/glog"
	"github.com/sezzle-calculator/gin"
	//"github.com/sezzle-calculator/gorm"
)

func main() {
	//Snag all flags that our application is run on.
	flag.Parse()
	flag.Lookup("alsologtostderr").Value.Set("true")

	//could not get go mysql to connect to my SQL server, so no db was used
	/*-------------------------------------------------------------------*/
	/*
	//Initalize our db.
	glog.Info("Initalizing db...")
	db, err := gorm.InitDB()
	if err != nil {
		glog.Fatal("Could not initalize db", err.Error())
	}

	//Defer this so that if our application exits, we close the db.
	//Double check this.

	defer db.Close()

	glog.Info("Initalizing Models...")

	err = gorm.Migrate()
	if err != nil {
		glog.Fatal("Could not run object migrations.")
	}
	*/
	/*-------------------------------------------------------------------*/

	gin.Run()

}
