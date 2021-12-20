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

import PatchStep from "postApi/steps/Patch";


async function updateChanges(step: StepT) {
    await PatchStep(step);
}


interface StepProps extends StepT{
    delStepCallback: (num: number) => Promise<void>
    updatable: boolean
}


const Step: React.FC<StepProps> = (props) => {
    const [active, change] = React.useState(false);
    const [title, changeTitle] = React.useState(props.title);
    const [description, changeDescr] = React.useState(props.description);

    return (
        <Box paddingBottom="15px">
            <VStack alignItems="left" minWidth="100px">
                <HStack width="100%" justifyContent="space-between">
                    {active && <Input height="30px" color="accent-3" borderColor="add-1"
                        value={title} placeholder='Введите название' 
                        onInput={(e) => {changeTitle(e.currentTarget.value)}}>
                    </Input>}

                    {!active && <Text textStyle="body" color="accent-3" height="30px" m="0px" 
                        paddingLeft="15px" paddingTop="0px">
                        {title}
                    </Text>}

                    {props.updatable && <HStack> 
                        <Button display="contents" onClick={() => 
                            {
                                active && updateChanges({recipe: props.recipe, num: props.num, title: title, description: description}); 
                                change(!active)}
                            }>

                            {!active && <PenIcon />}
                            {active && <OkIcon/>}
                        </Button> 

                        <Button display="contents" 
                            onClick={() => props.delStepCallback(props.num)}>
                            <DeleteIcon /> 
                        </Button>
                    </HStack>}
                </HStack>

                {active && <Textarea color="text" borderColor="add-1" boxSizing="content-box" 
                    value={description} placeholder="Введите описание"
                    onChange={(e) => {changeDescr(e.currentTarget.value)}}></Textarea>}

                {!active && <RoundBox display="inline-block" margin="5px" position="relative" borderColor="add-1"
                        width="100%" minHeight="70px" padding="3px" paddingLeft="10px"> 
                    <Text textStyle="text" color="body" height="100%" 
                        style={{margin: "0"}}>
                        {description}
                    </Text>
                </RoundBox>}
            </VStack>
        </Box>
    )
}

export default Step;