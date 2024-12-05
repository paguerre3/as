To publish your Docker image to Docker Hub, follow these steps:

### 1. **Login to Docker Hub from Your Terminal**
You need to log in to Docker Hub through the command line interface (CLI). Use the following command to log in:

```bash
docker login
```

You’ll be prompted to enter your Docker Hub username and password.

### 2. **Tag Your Docker Image IF NOT "already done" in Docker Compose**
Before you can push an image to Docker Hub, you need to tag it with your Docker Hub username and repository name. The tag follows the format:

```
<your-docker-username>/<repository-name>:<tag>
```

For example, if your Docker Hub username is `pablo` and you want to name your repository `my-go-app`, tag the image like this:

```bash
docker tag my-go-app pablo/my-go-app:latest
```

If you don’t specify a tag, Docker will automatically assume `latest` as the default tag.

### 3. **Push the Docker Image to Docker Hub**
Now, you can push the tagged image to Docker Hub using the `docker push` command:

```bash
docker push pablo/my-go-app:latest
```

This will upload your Docker image to Docker Hub. The upload process may take some time, depending on the size of your image and your internet connection.

### 4. **Verify the Image on Docker Hub**
Once the push is complete, go to your Docker Hub account, and you should see the image in your repositories list under the name `my-go-app`.

### 5. **Pull the Docker Image (to verify)**
To verify the image has been successfully pushed, you can pull the image from Docker Hub on any machine using the following command:

```bash
docker pull pablo/my-go-app:latest
```

This will pull the image to your local machine, confirming it's available on Docker Hub.

### Additional Notes:
- **Repository Visibility**: Make sure your Docker Hub repository is set to **Public** if you want others to access it. By default, repositories are private for new users.
- **Pushing Large Images**: If your image is large, Docker will use layers to upload the image in parts. This is normal.



---

### 1. **GitHub Codespaces (Free for Personal Use)**
- **Steps**:
  1. Push your Docker project to a GitHub repository.
  2. Use GitHub Codespaces to run and expose your container.
  3. Run:
     ```bash
     docker pull <your-dockerhub-username>/<image-name>
     docker run -d -p 8080:8080 <your-image-name>
     ```
     **OR build image form liked repository in GitHub Codespaces.**
     ```bash
     docker-compose -f docker-compose.yml up -d
     ```
  4. **Expose the Port Publicly**
    Expose the Port:

    In the top bar of your Codespaces IDE, click Ports.
    You’ll see a list of exposed ports. If the port (e.g., 8080) isn’t listed, click Add Port and enter the exposed port.
    Make the Port Public:

    Click the gear icon next to your port and select Port Visibility > Public.
    Get the Public URL:

    After making the port public, Codespaces will provide you with a public URL that you can use to access your service.

- **Limitations**: Limited hours per month on the free tier.



---
# Case: Damaged Spaceship 2

Trama:

Un suspiro de alivio escapa de tus labios al ver al robot reparador acoplarse a tu nave. La esperanza se renueva, pero dura poco. Una alarma estridente te saca de tu momentánea tranquilidad. El robot ha detectado una avería crítica: datos corruptos relacionados con la curva de "saturación y cambio de fase P-v" del fluido hidráulico. Sin esta información, la nave no puede calibrar sus actuadores y sigue a la deriva.

Una oleada de frustración te invade. ¡Tú eres un programador, no un ingeniero mecánico! Pero la desesperación da paso a la determinación. Siempre has sido bueno resolviendo problemas, y este no será la excepción.

La documentación del robot te da una pista: realizará 10 peticiones HTTP a la ruta /phase-change-diagram para intentar reconstruir el archivo corrupto. Ahí está tu oportunidad.

Pista:

Mientras buscas frenéticamente entre los manuales de la nave, encuentras el cuaderno de bitácora del ingeniero mecánico. La última entrada termina abruptamente con un "¡Wubba Lubba Dub-Dub!" garabateado y una mancha de lo que sospechas es salsa Sichuan...¡Pero entre diagramas a medio terminar y ecuaciones a medio resolver, encuentras la curva de saturación del fluido hidráulico!

PV diagram-----------------------------------------------------------------------[start diagram]

   P[MPa]    .                    .
     |       .                    .
     |       .          X=critical_point=pressure_const:PC=10MPa, temperature_const:Tc=500°C, volume_const:vc= 0.0035 m^3/kg 
     |       .         / \        .
     |       saturated/   \       .
     |   liquid line=/     \      .
     |       .      /       \     .
     |       .     /         \    .
     |       .    /           \   .
     |       .   /    saturated\  .
     |       . vapor (gas) line=\ .
     |       . /                 \.
0.05 |...……………/…………………………………………………\……………….......... T1const = 30°C (from 0.00105°C to 30°C (volume:V[m^3/kg]) at 0.05 pressure(P[MPa]))
     |       .                    .
     |       .                    .
     --------------------------------------------V [m^3/kg]
             .                    .
          0.00105               30.00 (30°C)

     "Repair robot will probe only T (temperature) > 30°C"   ¡Wubba Lubba Dub-Dub!

  Empirical constant for saturated vapor line is 30°C (from 0.00105°C to 30°C (volume:V[m^3/kg]) at 0.05 pressure(P[MPa])) and saturated liquid line starts at 0.00105°C while saturated vapor line starts at 30.00 30°C when both crosses at 0.05 pressure(P[MPa])).
  Then this is where they both saturated lines meet togheter critical_point=pressure_const:PC=10MPa, temperature_const:Tc=500°C, volume_const:vc= 0.0035 m^3/kg. 

---------------------------------------------------------------------------------[end diagram]


Ejemplo de una llamada del robot:

Ruta: [GET] /phase-change-diagram?pressure=10

Parámetros:

pressure: 10 (en mega pascales)

Respuesta Esperada:

```json
{
  "specific_volume_liquid": 0.0035,
  "specific_volume_vapor": 0.0035
}
```

Al registrar la URL de tu API, asegúrate de que sea accesible desde el exterior, tendrás tan solo 3 intentos y 5 minutos para que el robot repare el sistema hidráulico (aquí demuestras tu atención al detalle).