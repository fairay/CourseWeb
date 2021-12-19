import {
    Box,
    HStack,
    Text,
    VStack,
    Button,
    Input,
  } from "@chakra-ui/react";
import React from "react";

import RoundBox from "components/base/RoundBox";
import {Step as StepT} from "types/Step";

import PenIcon from "components/icon/Pen";
import DeleteIcon from "components/icon/DeleteSmall";


interface StepProps extends StepT{
    delStepCallback: (num: number) => Promise<void>
}


const Step: React.FC<StepProps> = (props) => {
    return (
        <Box paddingBottom="15px">
            <VStack alignItems="left" minWidth="100px">
                <HStack width="100%" justifyContent="space-between">
                    <Text textStyle="body" color="accent-3" height="30px" m="0px" 
                        paddingLeft="10px" paddingTop="5px">
                        {props.title}
                    </Text>

                    <HStack> 
                        <Button display="contents" 
                                /*onClick={() => props.delStepCallback(props.num)}*/>
                            <PenIcon />
                        </Button> 

                        <Button display="contents" 
                            onClick={() => props.delStepCallback(props.num)}>
                            <DeleteIcon /> 
                        </Button>
                    </HStack>
                </HStack>

                <RoundBox display="inline-block" margin="5px" position="relative" borderColor="add-1"
                        width="100%" minHeight="70px" padding="3px" paddingLeft="10px"> 
                    <Text textStyle="text" color="body" height="100%"
                        style={{margin: "0"}}>
                        {props.description}
                    </Text>
                </RoundBox>
            </VStack>
        </Box>
    )
}

export default Step;