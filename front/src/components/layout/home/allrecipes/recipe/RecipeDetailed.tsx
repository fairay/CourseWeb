import {
    Box,
    Button,
    Heading,
    HStack,
    Image,
    Link,
    Text,
    VStack,
  } from "@chakra-ui/react";
import React, { useState } from "react";

import photoRecipe from "img/photoRecipe.png";
import DeleteIcon from "components/icon/Delete";
import EmptyLike from "components/icon/EmptyLike";
import FullLike from "components/icon/FullLike";
import ClockIcon from "components/icon/Clock";
import PlateIcon from "components/icon/Plate";

import RoundBox from "components/base/RoundBox";


interface RecipeDtlProps {}

const RecipeDtl: React.FC<RecipeDtlProps> = (props) => {
    const [hide, show] = React.useState(true);

    return (
        <Box display="flex" w="100%"></Box>
    )
}

export default RecipeDtl;