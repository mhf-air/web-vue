#!/bin/sh

enable() {
  echo "abc"
}

disable() {
  echo "abc"
}

echo "enter e(enable) or d(disable)"
while :
do
  read action

  case $action in
    e)
      enable
      echo "port 9000 enabled"
      break
      ;;
  
    d)
      disable
      echo "port 9000 disabled"
      break
      ;;
  
    *)
      echo "Please enter a valid option"
      ;;
  esac
done
