import {
  Button,
  List,
  Table,
  TableCaption,
  Tbody,
  Td,
  Text,
  Tfoot,
  Th,
  Thead,
  Tr,
} from "@chakra-ui/react";
import React, { useState } from "react";
import IngredientItem from "./IngredientItem";
import Ingredient from "../ingredient/Ingredient"

interface IngredientsListProps {}

const IngredientsList: React.FC<IngredientsListProps> = (props) => {
  const [items, setItems] = useState([
    {
      id: 1,
      name: "Tomato",
      type: "vegetable",
      quantity: {
        num: 1,
        unit: "Kg",
      },
    },
    {
      id: 2,
      name: "Potato",
      type: "vegetable",
      quantity: {
        num: 2,
        unit: "Kg",
      },
    },
    {
      id: 3,
      name: "Onion",
      type: "vegetable",
      quantity: {
        num: 3,
        unit: "Kg",
      },
    },
  ]);

  const deleteItem = (id: number): void => {
    let newItems = items.filter((item) => item.id !== id);
    setItems(newItems);
  };
  return (
    <>
      <Table variant="simple" colorScheme="black" mt={2}>
        <Thead>
          <Tr>
            <Th fontSize="md">Ingredient</Th>
            <Th fontSize="md">Type</Th>
            <Th fontSize="md">Quantity</Th>
            <Th fontSize="md">Delete</Th>
          </Tr>
        </Thead>
        <Tbody>
          {items.map((item) => (
            <IngredientItem
              ingredient={item}
              key={item.id}
              deleteItem={deleteItem}
            />
          ))}
        </Tbody>
      </Table>

      <Ingredient />
      

      <Button colorScheme="green" d="block" my={6} w="100%">
        <Text>Search</Text>
      </Button>
    </>
  );
};

export default IngredientsList;
