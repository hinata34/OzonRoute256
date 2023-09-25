package commands

import (
	"bufio"
	"context"
	"fmt"
	"homework-5/internal/app/address"
	database "homework-5/internal/app/db"
	"homework-5/internal/app/user"
	"os"
	"strconv"
	"strings"
)

type DBService struct{}

func NewDBService() *DBService {
	return &DBService{}
}

func (db *DBService) Run(ctx context.Context, args []string) error {
	database, err := database.NewDB(ctx)
	if err != nil {
		return err
	}
	defer database.GetPool(ctx).Close()

	usersRepo := user.NewUsers(database)
	addressesRepo := address.NewAddresses(database)

	DBInput(ctx, usersRepo, addressesRepo)

	return nil
}

func (db *DBService) Description() string {
	return "Connect to users and addresses Database"
}

func DBInput(ctx context.Context, usersRepo *user.UsersRepo, addressesRepo *address.AddressesRepo) {
	reader := bufio.NewReader(os.Stdin)
	exit := false
	for {
		fmt.Println("Choose Table: users, addresses:")
		fmt.Print("->")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Program Exit")
			return
		}
		text = strings.Replace(text, "\n", "", -1)
		table := strings.Split(text, " ")

		fmt.Println("Input command: Add, GetById, List, Update, Exit:")
		fmt.Print("->")
		text, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Program Exit")
			return
		}
		text = strings.Replace(text, "\n", "", -1)
		query := strings.Split(text, " ")

		switch table[0] {
		case "users":
			exit, err = tableUsers(ctx, usersRepo, query)
		case "addresses":
			exit, err = tableAddresses(ctx, addressesRepo, query)
		default:
			fmt.Println("Wrong Table")
		}

		if err != nil {
			fmt.Println("Program Exit")
			return
		}

		if exit {
			fmt.Println("Program Exit Correctly")
			break
		}
	}
}

func tableUsers(ctx context.Context, usersRepo *user.UsersRepo, query []string) (bool, error) {
	switch query[0] {
	case "Add":
		age, err := strconv.ParseInt(query[2], 10, 64)
		if err != nil {
			return true, err
		}
		user := &user.User{Name: query[1], Age: age}
		id, err := usersRepo.Add(ctx, user)
		if err != nil {
			return true, err
		}
		fmt.Printf("Added %d\n", id)
	case "GetById":
		id, err := strconv.ParseInt(query[1], 10, 64)
		if err != nil {
			return true, err
		}
		user, err := usersRepo.GetById(ctx, id)
		if err != nil {
			return true, err
		}
		fmt.Printf("Got User: id=%d, name=%s, age=%d\n", user.ID, user.Name, user.Age)
	case "List":
		users, err := usersRepo.List(ctx)
		if err != nil {
			return true, err
		}
		for _, user := range users {
			fmt.Printf("Got User: id=%d, name=%s, age=%d\n", user.ID, user.Name, user.Age)
		}
	case "Update":
		id, err := strconv.ParseInt(query[1], 10, 64)
		if err != nil {
			return true, err
		}
		age, err := strconv.ParseInt(query[3], 10, 64)
		if err != nil {
			return true, err
		}
		user := &user.User{ID: id, Name: query[2], Age: age}
		res, err := usersRepo.Update(ctx, user)
		if err != nil {
			return true, err
		}
		if res {
			fmt.Println("User Updated Successfully")
		} else {
			fmt.Println("User Updated Failed")
		}
	case "Exit":
		return true, nil
	default:
		fmt.Println("Wrong command\nUse: Add, GetById, List, Update, Exit")
	}
	return false, nil
}

func tableAddresses(ctx context.Context, addressesRepo *address.AddressesRepo, query []string) (bool, error) {
	switch query[0] {
	case "Add":
		house_number, err := strconv.ParseInt(query[1], 10, 64)
		if err != nil {
			return true, err
		}

		user_id, err := strconv.ParseInt(query[3], 10, 64)
		if err != nil {
			return true, err
		}

		address := &address.Address{HouseNumber: house_number, StreetName: query[2], UserID: user_id}
		id, err := addressesRepo.Add(ctx, address)
		if err != nil {
			return true, err
		}
		fmt.Printf("Added %d\n", id)
	case "GetById":
		id, err := strconv.ParseInt(query[1], 10, 64)
		if err != nil {
			return true, err
		}
		address, err := addressesRepo.GetById(ctx, id)
		if err != nil {
			return true, err
		}
		fmt.Printf("Got User: id=%d, house_number=%d, street_name=%s, user_id=%d\n", address.ID, address.HouseNumber, address.StreetName, address.UserID)
	case "List":
		addresses, err := addressesRepo.List(ctx)
		if err != nil {
			return true, err
		}
		for _, address := range addresses {
			fmt.Printf("Got User: id=%d, house_number=%d, street_name=%s, user_id=%d\n", address.ID, address.HouseNumber, address.StreetName, address.UserID)
		}
	case "Update":
		id, err := strconv.ParseInt(query[1], 10, 64)
		if err != nil {
			return true, err
		}

		house_number, err := strconv.ParseInt(query[2], 10, 64)
		if err != nil {
			return true, err
		}

		user_id, err := strconv.ParseInt(query[4], 10, 64)
		if err != nil {
			return true, err
		}

		address := &address.Address{ID: id, HouseNumber: house_number, StreetName: query[3], UserID: user_id}

		res, err := addressesRepo.Update(ctx, address)
		if err != nil {
			return true, err
		}
		if res {
			fmt.Println("User Updated Successfully")
		} else {
			fmt.Println("User Updated Failed")
		}
	case "Exit":
		return true, nil
	default:
		fmt.Println("Wrong command\nUse: Add, GetById, List, Update, Exit")
	}
	return false, nil
}
