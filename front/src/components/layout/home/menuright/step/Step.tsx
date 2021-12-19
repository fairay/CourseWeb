import {
    Box,
    HStack,
    Text,
    VStack,
    Button,
  } from "@chakra-ui/react";
import React from "react";

import RoundBox from "components/base/RoundBox";
import Input from "components/base/Input";
import Textarea from "components/base/TextArea";
import {Step as StepT} from "types/Step";

import PenIcon from "components/icon/Pen";
import DeleteIcon from "components/icon/DeleteSmall";
import OkIcon from "components/icon/Ok";


interface StepProps extends StepT{
    delStepCallback: (num: number) => Promise<void>
}


const Step: React.FC<StepProps> = (props) => {
    const [active, change] = React.useState(false);

    return (
        <Box paddingBottom="15px">
            <VStack alignItems="left" minWidth="100px">
                <HStack width="100%" justifyContent="space-between">
                    {active && <Input height="30px" color="accent-3" borderColor="add-1"
                        value={props.title} placeholder="Введите название" 
                        onInput={(e) => {props.title = e.currentTarget.value}}>
                    </Input>}

                    {!active && <Text textStyle="body" color="accent-3" height="30px" m="0px" 
                        paddingLeft="15px" paddingTop="0px">
                        {props.title}
                    </Text>}

                    <HStack> 
                        <Button display="contents" onClick={() => change(!active)}>
                            {!active && <PenIcon />}
                            {active && <OkIcon/>}
                        </Button> 

                        <Button display="contents" 
                            onClick={() => props.delStepCallback(props.num)}>
                            <DeleteIcon /> 
                        </Button>
                    </HStack>
                </HStack>

                {active && <Input color="text" borderColor="add-1"
                    value={props.description} placeholder="Введите описание"
                    onInput={(e) => {props.description = e.currentTarget.value}}></Input>}

                {!active && <RoundBox display="inline-block" margin="5px" position="relative" borderColor="add-1"
                        width="100%" minHeight="70px" padding="3px" paddingLeft="10px"> 
                    <Text textStyle="text" color="body" height="100%"
                        style={{margin: "0"}}>
                        {props.description}
                    </Text>
                </RoundBox>}
            </VStack>
        </Box>
    )
}

export default Step;