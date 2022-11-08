package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/IMDb-searcher/internal/cli/formatter"
	"github.com/IMDb-searcher/internal/logger"
)

type IUserInterface interface {
	Run()
}

type UserInterface struct {
	IUserInterface

	formatter formatter.IFormatter
	log       logger.ILogger
}

func getCommandLineInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func (u *UserInterface) choiceWrapper(requestHandler func(string) (string, error)) {
	request := getCommandLineInput()

	responce, err := requestHandler(request)
	if err != nil {
		u.log.Error(err)
		fmt.Println(err)
	} else {
		u.log.Info(responce)
		fmt.Println(responce)
	}
}

func (u *UserInterface) Run() {
	for {
		fmt.Println()
		fmt.Println()
		fmt.Println("Please make your choice (0-4):")
		fmt.Println("0. Exit")
		fmt.Println("1. Find info by person name")
		fmt.Println("2. Find title and cast info by title name")
		fmt.Println("3. Find titles by person name")
		fmt.Println("4. Find all titles by specific year")

		choice := getCommandLineInput()
		u.log.Info(choice)

		switch choice {
		case "0":
			fmt.Println("Bye!")
			return
		case "1":
			fmt.Println("Enter the name of the searched person:")
			u.choiceWrapper(u.formatter.FindInfoByPersonName)
		case "2":
			fmt.Println("Enter the name of the searched title:")
			u.choiceWrapper(u.formatter.FindTitleAndCastInfoByTitleName)
		case "3":
			fmt.Println("Enter the name of the person for whom you want to search for titles:")
			u.choiceWrapper(u.formatter.FindTitlesByPersonName)
		case "4":
			fmt.Println("Enter the year for whom you want to search for titles:")
			u.choiceWrapper(u.formatter.FindAllTitlesBySpecificYear)
		default:
			fmt.Println("Input error. Please make sure that the selection contains only numbers from 0 to 4.")
		}
	}

}

func New(logger logger.ILogger) (IUserInterface, error) {
	fmt.Println("The database is loading...") // TO DO: Need to make better

	formatter, err := formatter.New(logger)
	if err != nil {
		return nil, err
	}
	return &UserInterface{
		formatter: formatter,
		log:       logger,
	}, nil
}
