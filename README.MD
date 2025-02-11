# Firebase Admin Manager Go (CLI Interactiva - En Desarrollo)

Una herramienta CLI en Go para administrar usuarios de Firebase, con enfoque en la gestión de roles (custom claims).  Actualmente en desarrollo para convertirla en una aplicación CLI interactiva.

## Descripción

Este proyecto proporciona una aplicación de línea de comandos (CLI) construida en Go para interactuar con un proyecto de Firebase y realizar tareas de administración de usuarios.  Utiliza el SDK de administrador de Firebase para Go.  La versión actual soporta comandos individuales, pero se está trabajando en una interfaz interactiva completa.

**Funcionalidad Actual (Comandos Individuales):**

*   **Listar usuarios:** Obtener una lista de usuarios, con opciones de paginación.
*   **Obtener detalles de un usuario:** Ver la información detallada de un usuario, incluyendo sus custom claims (roles).
*   **Añadir roles:** Agregar roles a un usuario.
*   **Establecer roles:** Establecer/reemplazar los roles de un usuario.
*   **Eliminar usuarios:** Eliminar un usuario por su UID.

**Funcionalidad Planificada (CLI Interactiva):**

La próxima versión de esta herramienta será una aplicación CLI interactiva. En lugar de usar comandos separados, el usuario ejecutará la aplicación y se le presentará un menú de opciones para realizar las diferentes tareas de administración de usuarios.  Esto hará que la herramienta sea más fácil de usar y más intuitiva.

## Requisitos

*   Go: Versión 1.22.3 o superior.
*   Firebase Project: Un proyecto de Firebase configurado.
*   Cuenta de Servicio:  Con permisos de administrador (rol `Firebase Admin SDK Administrator Service Agent`). Archivo JSON (`serviceAccountKey.json`).
*   Variable de entorno: `GOOGLE_APPLICATION_CREDENTIALS` (ruta absoluta al `serviceAccountKey.json`).  O usar un archivo `.env`.

## Instalación

1.  **Clonar el repositorio:**

    ```bash
    git clone <URL_DE_TU_REPOSITORIO>
    cd <NOMBRE_DEL_DIRECTORIO>
    ```

2.  **Dependencias:**

    ```bash
    go mod tidy
    ```

3.  **Credenciales de Firebase:**
    * Descarga `serviceAccountKey.json` (Consola de Firebase -> Configuración del proyecto -> Cuentas de servicio -> Generar nueva clave privada). *¡No lo incluyas en Git!*

    
## Uso (Versión Actual - Comandos Individuales)

```bash
# Compilar (opcional, pero recomendado)
go build -o firebase-admin ./cmd/main.go

# O usar go run
go run ./cmd/main.go <comando> [argumentos]
