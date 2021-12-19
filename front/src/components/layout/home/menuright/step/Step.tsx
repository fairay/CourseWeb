import {
    Box,
    Text,
    VStack,
  } from "@chakra-ui/react";
import React from "react";

import RoundBox from "components/base/RoundBox";
import {Step as StepT} from "types/Step";

interface StepProps extends StepT{}

const Step: React.FC<StepProps> = (props) => {
    return (
        <Box paddingBottom="15px">
            <VStack alignItems="left" minWidth="100px">
                <Text textStyle="body" color="accent-3" height="30px" m="0px" 
                    paddingLeft="10px" paddingTop="5px">
                    {props.title}
                </Text>

                <RoundBox display="inline-block" margin="5px" position="relative" borderColor="add-1"
                        width="100%" minHeight="70px" padding="3px" paddingLeft="10px"> 
                    <Text textStyle="text" color="body" height="20px"
                        style={{margin: "0"}}>
                        {props.description}
                    </Text>
                </RoundBox>
            </VStack>
        </Box>
    )
}

export default Step;