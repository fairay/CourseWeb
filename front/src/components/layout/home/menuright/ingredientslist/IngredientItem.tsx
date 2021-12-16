import { CloseIcon } from "@chakra-ui/icons";
import { Button, IconButton, ListItem, Td, Text, Tr } from "@chakra-ui/react";
import React from "react";

interface IngredientItemProps {
  ingredient: {
    id: number;
    name: string;
    type: string;
    quantity: {
      num: number;
      unit: string;
    };
  };
  deleteItem: (id: number) => void;
}

const IngredientItem: React.FC<IngredientItemProps> = ({
  ingredient,
  deleteItem,
}) => {
  const {
    id,
    name,
    quantity: { num, unit },
    type,
  } = ingredient;

  return (
    <Tr>
      <Td>{name}</Td>
      <Td>{type}</Td>
      <Td>{`${num} ${unit}`}</Td>
      <Td>
        <IconButton
          aria-label="Delete ingredient"
          colorScheme="red"
          icon={<CloseIcon h={3} />}
          onClick={() => deleteItem(id)}
        />
      </Td>
    </Tr>
  );
};

export default IngredientItem;
