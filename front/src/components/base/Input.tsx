import React from "react";
import {
    InputProps as IProps, Input as I
} from "@chakra-ui/react";
import RoundBox from "./RoundBox"
import theme from "extendTheme"

interface InputProps extends IProps {
    name?: string
    type?: string
    placeholder?: string
    focusBorderColor?: string
}

const Input: React.FC<InputProps> = (props) => {
    const {
        name = "", type = "",
        placeholder = "Введите данные",
        focusBorderColor = theme.colors["accent-3"],
         ...rest 
     } = props

     const style = {
        color: '#000000',
        boxShadow: '0 0 5px 1px ' + focusBorderColor,
      };

    return (
    <RoundBox {...rest}>
        <I 
            name={name} type={type}
            borderWidth={0} borderRadius={10} 
            width="100%"
            placeholder={placeholder}
            _focus={style}
        />
    </RoundBox>
    )
}

export default Input