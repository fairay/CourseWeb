import {
  Box,
  Button,
  Input,
  InputGroup,
  InputRightAddon,
  Menu,
  MenuButton,
  MenuItemOption,
  MenuList,
  MenuOptionGroup,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Text,
  useDisclosure,
} from "@chakra-ui/react";
import React from "react";

interface IngredientsFormProps {}

const IngredientsForm: React.FC<IngredientsFormProps> = ({}) => {
  const { isOpen, onOpen, onClose } = useDisclosure();
  return (
    <Box w="100%" mt={2}>
      <Button
        d="block"
        w="100%"
        border="2px solid"
        borderColor="primary"
        backgroundColor="third"
        colorScheme="yellow"
        onClick={onOpen}
      >
        <Text color="primary-dark">Add an ingredient</Text>
      </Button>
      <Modal
        isCentered
        onClose={onClose}
        isOpen={isOpen}
        motionPreset="slideInBottom"
      >
        <ModalOverlay />
        <ModalContent backgroundColor="secondary">
          <ModalHeader>Add an ingredient</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Text my={2}>Type :</Text>
            <Input
              backgroundColor="white"
              placeholder="What type is your ingredient"
              size="sm"
            />
            <Text my={2}>Name :</Text>
            <Input
              backgroundColor="white"
              placeholder="What is the name of your ingredient"
              size="sm"
            />
            <Text my={2}>Quantity :</Text>
            <InputGroup size="sm">
              <Input
                backgroundColor="white"
                placeholder="What how much you have of this ingredient"
              />
              <InputRightAddon
                children={
                  <Menu closeOnSelect={false}>
                    <MenuButton colorScheme="blue">Unit : Kg</MenuButton>
                    <MenuList minWidth="240px">
                      <MenuOptionGroup defaultValue="kg" type="radio">
                        <MenuItemOption value="kg">Kg</MenuItemOption>
                        <MenuItemOption value="oz">Oz</MenuItemOption>
                        <MenuItemOption value="cups">Cups</MenuItemOption>
                      </MenuOptionGroup>
                    </MenuList>
                  </Menu>
                }
              />
            </InputGroup>
          </ModalBody>
          <ModalFooter>
            <Button
              border="2px solid"
              borderColor="primary"
              backgroundColor="third"
              colorScheme="yellow"
              mr={3}
              onClick={onClose}
            >
              <Text color="primary-dark">Add</Text>
            </Button>{" "}
            <Button mr={3} colorScheme="red" onClick={onClose}>
              Cancel
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </Box>
  );
};

export default IngredientsForm;
