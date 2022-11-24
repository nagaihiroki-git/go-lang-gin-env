import React from 'react'
import {IconContext} from 'react-icons'
import {BsFillCartFill} from 'react-icons/all'

export const Cart = () => {
    return (
        <>
            <IconContext.Provider value={{size: '2em'}}>
            <BsFillCartFill/>
                カート
            </IconContext.Provider>
        </>
    )
}