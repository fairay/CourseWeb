import { Box, Button, useDisclosure } from '@chakra-ui/react'
import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
} from '@chakra-ui/react'
import React from 'react'

import AddIcon from 'components/icon/Add'
import Input from 'components/base/Input'
import { Ingredient as IngredientT } from 'types/Ingredient'

export default function IngredientModal(propos) {
  const { isOpen, onOpen, onClose } = useDisclosure()
  var data: IngredientT = { title: '', amount: '' }

  async function put() {
    await propos.putCallback(data)
    onClose()
  }

  return (
    <>
      <Button display="contents" onClick={onOpen}>
        <Box paddingLeft="15px"> <AddIcon /> </Box>
      </Button>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent bg='bg'>
          <ModalHeader>Новый ингредиент</ModalHeader>
          <ModalCloseButton />
          <ModalBody columnGap="6px" display='flex' flexDirection='row' alignItems="center">
            <Box rowGap="6px" display='flex' flexDirection='column' width="100%">
              <Input value={data.title} placeholder='Введите ингредиент'
                onInput={(e) => {data.title = e.currentTarget.value}}
              />
              <Input value={data.amount} placeholder='Введите меру' 
                textColor='accent-3' placeholderColor='accent-3'
                onInput={(e) => {data.amount = e.currentTarget.value}}
              />
            </Box>
            <Button display="contents" onClick={put}>
              <AddIcon width='50' height='40' />
            </Button>
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  )
}
