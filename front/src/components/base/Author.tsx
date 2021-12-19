import React from "react";
import {
    Box, BoxProps
} from "@chakra-ui/react";
import RoundBox, { RoundBoxProps } from "./RoundBox";
import AuthorIcon from "components/icon/Author"

interface AuthorProps extends RoundBoxProps {
    authorName?: string
}

const Author: React.FC<AuthorProps> = (props) => {
    const { authorName="god saves GSPD", ...rest } = props
    
    return (
        <RoundBox {...rest}>
            <Box marginY="10px" marginRight="10px" marginLeft="12px">
                <AuthorIcon />
            </Box>
            <Box id="author" width={"100%"} display="flex" alignItems="center" textStyle="caption">
                {authorName}
            </Box>
        </RoundBox>
    )
}

export default Author