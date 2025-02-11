package firebase

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth" // Import sin alias
	"google.golang.org/api/iterator"

	"google.golang.org/api/option"
)

// FirebaseApp representa la conexión con Firebase.
type FirebaseApp struct {
	Auth *auth.Client
	Ctx  context.Context
}

// NewFirebaseApp inicializa la conexión con Firebase.
func NewFirebaseApp() (*FirebaseApp, error) {
	ctx := context.Background()

	saPath := "/home/yeison/Documentos/Mis proyectos/billar/firebase-admin-manager-go/service-account.json"
	if saPath == "" {
		return nil, fmt.Errorf("GOOGLE_APPLICATION_CREDENTIALS no esta definida")
	}
	opt := option.WithCredentialsFile(saPath)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}

	return &FirebaseApp{
		Auth: authClient,
		Ctx:  ctx,
	}, nil
}

// ListUsers lista los usuarios de Firebase.
func (f *FirebaseApp) ListUsers(pageSize int) ([]*auth.ExportedUserRecord, error) { // <-- CORREGIDO
	var users []*auth.ExportedUserRecord // <-- CORREGIDO
	iter := f.Auth.Users(f.Ctx, "")      // Inicia la iteración. "" significa desde el principio.
	for {
		user, err := iter.Next() // ¡Sin argumentos!
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error listing users: %v\n", err)
			return nil, err
		}
		users = append(users, user)

		// Control de tamaño de página (manual, pero efectivo)
		if len(users) >= pageSize {
			break
		}
	}
	return users, nil
}

// ListAllUsers lista *todos* los usuarios de Firebase (sin límite de tamaño de página).
func (f *FirebaseApp) ListAllUsers() ([]*auth.ExportedUserRecord, error) { // <-- CORREGIDO
	var users []*auth.ExportedUserRecord // <-- CORREGIDO
	iter := f.Auth.Users(f.Ctx, "")
	for {
		user, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error listing users: %v", err)
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUser obtiene un usuario por UID.
func (f *FirebaseApp) GetUser(uid string) (*auth.UserRecord, error) {
	user, err := f.Auth.GetUser(f.Ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("error getting user %s: %v", uid, err)
	}
	return user, nil
}

// SetUserRoles establece los roles (custom claims) de un usuario.
func (f *FirebaseApp) SetUserRoles(uid string, roles []string) error {
	claims := map[string]interface{}{
		"roles": roles,
	}
	err := f.Auth.SetCustomUserClaims(f.Ctx, uid, claims)
	if err != nil {
		return fmt.Errorf("error setting custom claims for user %s: %v", uid, err)
	}
	return nil
}

// DeleteUser elimina un usuario por UID.
func (f *FirebaseApp) DeleteUser(uid string) error {
	err := f.Auth.DeleteUser(f.Ctx, uid)
	if err != nil {
		return fmt.Errorf("error deleting user %s: %v", uid, err)
	}
	return nil
}
