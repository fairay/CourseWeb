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
        <Box>
            <VStack alignItems="left" marginLeft="10px" marginRight="20px" minWidth="100px">
                <Text textStyle="body" color="accent-3" height="30px" m="0px">
                    {props.title}
                </Text>

                <RoundBox display="inline-block" margin="5px" position="relative" borderColor="add-1"
                        width="100%" minHeight="70px" padding="3px"> 
                    <Text textStyle="caption" color="accent-3" height="20px"
                        style={{margin: "0"}}>
                        {props.description}
                    </Text>
                </RoundBox>
            </VStack>
        </Box>
    )
}

export default Step;