import React from "react";
import {
    InputProps as IProps, Input as I,
    Box, Text
} from "@chakra-ui/react";
import RoundBox from "../RoundBox";
import theme from "extendTheme"
import ClockIcon from "components/icon/Clock";

interface InputProps extends IProps {
    duration?: number
}

const ClockBox: React.FC<InputProps> = (props) => {

      var stringDuration = ""

        if (!props.duration)
            stringDuration = "---"
        else if (props.duration < 90)
            stringDuration += props.duration + " мин"
        else if (props.duration < 60 * 24)
            stringDuration += Math.round(props.duration / 60) + " ч"
        else
            stringDuration += Math.round(props.duration / (60 * 24)) + " д"

    return (
    <RoundBox width="120px" height="30px" 
        borderColor="accent-1" alignItems="center"> 
            <Box marginY="auto" marginX="8px"> <ClockIcon /> </Box>
            <Text marginLeft="2px" textStyle="body"> {stringDuration}  </Text>
    </RoundBox>
    )
}

export default ClockBox