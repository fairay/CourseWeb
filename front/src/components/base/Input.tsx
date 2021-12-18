import React from "react";
import {
    BoxProps, Input as I
} from "@chakra-ui/react";
import RoundBox from "./RoundBox"
import theme from "extendTheme"

interface InputProps extends BoxProps {
    placeholder?: string
    focusBorderColor?: string
}

const Input: React.FC<InputProps> = (props) => {
    const {
        placeholder = "Введите данные",
        focusBorderColor = theme.colors["accent-1"],
         ...rest 
     } = props

     const style = {
        color: '#000000',
        'box-shadow': '0 0 5px 1px ' + focusBorderColor,
      };

    return (
    <RoundBox {...rest}>
        <I 
            borderWidth={0} borderRadius={10} 
            width="100%"
            placeholder={placeholder}
            _focus={style}
        />
    </RoundBox>
    )
}

export default Input