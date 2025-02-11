package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"firebase.google.com/go/v4/auth"
	fbadmin "github.com/Yeison07/firebase-admin-manager-go/internal"
	"github.com/spf13/cobra"
)

func main() {
	app, err := fbadmin.NewFirebaseApp()
	if err != nil {
		log.Fatal(err)
	}

	var rootCmd = &cobra.Command{Use: "firebase-admin"}

	// Comando para listar usuarios
	var listUsersCmd = &cobra.Command{
		Use:   "list-users",
		Short: "Lista los usuarios de Firebase",
		Run: func(cmd *cobra.Command, args []string) {
			pageSize, _ := cmd.Flags().GetInt("page-size")
			if pageSize <= 0 {
				pageSize = 100 // Valor predeterminado si no se especifica o es inválido
			}

			// Usar ListAllUsers si pageSize es muy grande (o un valor especial, como -1)
			var users []*auth.ExportedUserRecord
			if pageSize > 1000 { // o if pageSize == -1, por ejemplo
				users, err = app.ListAllUsers()
			} else {
				users, err = app.ListUsers(pageSize)
			}

			if err != nil {
				log.Fatal(err)
			}
			for _, user := range users {
				fmt.Printf("UID: %s, Email: %s, Roles: %v\n", user.UID, user.Email, user.CustomClaims["roles"])
			}

		},
	}
	listUsersCmd.Flags().Int("page-size", 100, "Tamaño de página para la lista de usuarios (máximo 1000, o un valor mayor para obtener todos)")
	rootCmd.AddCommand(listUsersCmd)

	// Comando para obtener un usuario
	var getUserCmd = &cobra.Command{
		Use:   "get-user <uid>",
		Short: "Obtiene un usuario por UID",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			uid := args[0]
			user, err := app.GetUser(uid)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("UID: %s\n", user.UID)
			fmt.Printf("Email: %s\n", user.Email)
			fmt.Printf("Display Name: %s\n", user.DisplayName)
			fmt.Printf("Email Verified: %t\n", user.EmailVerified)
			fmt.Printf("Disabled: %t\n", user.Disabled)
			fmt.Printf("Roles: %v\n", user.CustomClaims["roles"])
		},
	}
	rootCmd.AddCommand(getUserCmd)

	// Comando para establecer roles
	var setRolesCmd = &cobra.Command{
		Use:   "set-roles <uid> <role1,role2,...>",
		Short: "Establece los roles de un usuario",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			uid := args[0]
			roles := strings.Split(args[1], ",")
			err := app.SetUserRoles(uid, roles)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Roles establecidos para el usuario %s\n", uid)
		},
	}
	rootCmd.AddCommand(setRolesCmd)

	// Comando para eliminar un usuario
	var deleteUserCmd = &cobra.Command{
		Use:   "delete-user <uid>",
		Short: "Elimina un usuario por UID",
		Args:  cobra.ExactArgs(1), // Asegúrate de que haya exactamente un argumento (el UID)
		Run: func(cmd *cobra.Command, args []string) {
			uid := args[0]
			err := app.DeleteUser(uid)
			if err != nil {
				log.Fatalf("Error deleting user: %v\n", err) // Mejor manejo de errores
				return
			}
			fmt.Printf("Usuario con UID %s eliminado exitosamente.\n", uid)
		},
	}
	rootCmd.AddCommand(deleteUserCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
