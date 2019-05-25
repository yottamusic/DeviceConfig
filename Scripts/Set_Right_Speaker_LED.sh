#!/bin/sh

if [ $# != 2 ]; then
  exit 0
fi

# parameter 1:
# speaker serial

# parameter 2: 
# 0x20 for red
# 0x21 for green
# 0x22 for blue

# parameter 3:
# bit 0-2:
#   1: fast blinking
#   2: slow blinking
#   3: breathing
#   4: on
#   5: off
# bit 7 will be polarity control

/root/i2crw0 w 33 $1 0x21 1


