import React from "react";
import {
    TextareaProps as IProps, Textarea as I
} from "@chakra-ui/react";
import RoundBox from "./RoundBox"
import theme from "extendTheme"

interface InputProps extends IProps {
    name?: string
    type?: string
    placeholder?: string
    focusBorderColor?: string
    placeholderColor?: string
}

const Textarea: React.FC<InputProps> = (props) => {
    const {
        name = "", type = "",
        placeholder = "Введите данные",
        textColor = theme.colors["text"],
        placeholderColor = theme.colors["text"],
        focusBorderColor = theme.colors["accent-3"],
        onChange = () => {},
         ...rest 
     } = props

     const style = {
        color: textColor,
        boxShadow: '0 0 5px 1px ' + focusBorderColor,
      };

    return (
    <RoundBox>
        <I 
            name={name} type={type}
            textColor={textColor}
            borderWidth={0} borderRadius={10} 
            width="100%"
            placeholder={placeholder}
            onChange={onChange}
            _focus={style}
            _placeholder={{
                color: placeholderColor,
                opacity: 0.6
            }}
            {...rest}
        />
    </RoundBox>
    )
}

export default Textarea