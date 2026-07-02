#!/bin/bash

LOGDIR=/var/log/deneme/

DATE=$(date +"%d-%m-%Y-%H-%M")

FILE_SIZE_LIMIT=$((10*1024*1024))

ROTATOR_LOGDIR=/var/log/rotator/

PROCESS_LOG="$ROTATOR_LOGDIR/rotator_process.log"

mkdir -p "$ROTATOR_LOGDIR"

# BU SATIRDAN SONRAKI tum ciktilari kalici log dosyasina yonlendiriyoruz 
exec >> "$PROCESS_LOG" 2>&1

for LOGFILE in $(find $LOGDIR -maxdepth 1 -type f -name '*.log'); do
    BASENAME=$(basename "$LOGFILE")
    # -f file demek dosyanın olup olmadığını checkliyor.
    if [ -f ${LOGFILE} ]; then
        FILESIZE=$(stat -c %s ${LOGFILE})
    # -gt = greater than demek.
        if [ ${FILESIZE} -gt ${FILE_SIZE_LIMIT} ]; then
            echo "[UYARI]: DOSYA BOYUTU $FILESIZE 10MB'DEN BUYUK!"
            echo "${LOGFILE} yedegi aliniliyor..."
            tar -czvf ${LOGFILE}_${DATE}.tar.gz -C ${LOGDIR} ${BASENAME}
            echo "Dosya basari ile sikistirildi: ${LOGFILE}_${DATE}.tar.gz"
            echo "${LOGFILE} temizleniyor..."
            #truncate islemi deniyor, acık dosyayı kapatmadan pid patlatmadan temiliyo 0 bayt yapıyo icerigini.
            > "$LOGFILE"
            echo "Log dosyasi basariyla temizlendi."
        else
        echo "Dosya boyutu normal aralikta: $FILESIZE"
        fi
    else
        echo "Dosya bulunamadi."
    fi
done
        find $LOGDIR -name "*.tar.gz" -mtime +30 -delete
        echo "30 gunden eski arsivler silindi."


