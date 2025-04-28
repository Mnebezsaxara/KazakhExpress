   # Проверка статуса Docker
   if (-not (docker info 2>$null)) {
       Write-Host "Docker не запущен. Запускаю Docker Desktop..."
       Start-Process "C:\Program Files\Docker\Docker\Docker Desktop.exe"
       Start-Sleep -Seconds 30  # Ждем запуска Docker
   }

   # Сборка и запуск через Docker
   docker-compose up --build -d

   # Проверка статуса
   docker-compose ps