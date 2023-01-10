# Email-app
## Crear archivo .env con las variables de entorno:
- EMAIL_HOST
- EMAIL_PORT
- EMAIL_SENDER_USER
- EMAIL_SENDER_PASSWORD

### Para configurar una contraseña de app en google (contraseña que utilizaremos para nuestra conexion con smtp de google):
1. Ve a tu cuenta de google.
2. Selecciona seguridad.
3. Verifica tener activada la autenticacion de 2 factores (en caso de no tenerla, activarla)
4. Ir a Iniciar sesion en Google: "contraseñas de aplicaciones"
5. Seleccionar aplicacion: Correo - Seleccionar Dispositivo: Otra
6. Se generara una contraseña de aplicacion, la cual utilizaremos en nuestro archivo .env para setear como variable de entorno (EMAIL_SENDER_PASSWORD)


