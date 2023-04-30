#!/usr/bin/env bash
#******* Global Variables *******
# shellcheck disable=SC2034
rosto_menu="https://www.rosto.sk/media/content/media/NOVE%2021/Mi%C5%A1ka/MENU%20AJ/1170x2083.jpg"
cool_bowling_menu="https://www.coolbowling.sk/-denne-menu" 
image_name="rosto-menu.jpg"
# shellcheck disable=SC2034
version="1.0"

#******* Functions *******

#Rosto
wget  https://www.rosto.sk/media/content/media/NOVE%2021/Mi%C5%A1ka/MENU%20AJ/1170x2083.jpg
mv 1170x2083.jpg ${image_name}


#CoolBouling
 wkhtmltoimage --crop-y 330 --crop-h 830 ${cool_bowling_menu}  ./cool-menu.png

#Pilsner-menu
wkhtmltoimage --crop-y 360 --crop-h 2770 https://www.pilsnerurquellpub.sk/kosice/denne-menu   ./pilsner-menu.png