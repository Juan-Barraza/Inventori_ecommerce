#!/bin/bash
cd /home/ubuntu/Inventori_ecommerce

# 1. Obtener el backup más reciente de código y DB
LATEST_CODE_BACKUP=$(ls -t backup_code_*.zip | head -n 1)
LATEST_DB_BACKUP=$(ls -t backup_*.sql | head -n 1)

if [ -z "$LATEST_CODE_BACKUP" ] || [ -z "$LATEST_DB_BACKUP" ]; then
    echo "❌ No se encontraron backups completos!"
    exit 1
fi

echo "🔁 Iniciando rollback..."
echo "📦 Código: $LATEST_CODE_BACKUP"
echo "🗃️ Base de datos: $LATEST_DB_BACKUP"

# 2. Detener contenedores
docker-compose down

# 3. Restaurar código
unzip -o "$LATEST_CODE_BACKUP" -d /home/ubuntu/Inventori_ecommerce
rsync -a /home/ubuntu/Inventori_ecommerce/ /home/ubuntu/Inventori_ecommerce/
rm -rf /home/ubuntu/Inventori_ecommerce

# 4. Restaurar DB
docker-compose -f docker-compose.yml -f docker-compose.prod.yml run --rm db \
    psql -U postgres -d mi_db -f "$LATEST_DB_BACKUP"

# 5. Reconstruir y levantar
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d --build

echo "✅ Rollback completado. Versión anterior restaurada."