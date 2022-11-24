import {IconContext} from 'react-icons'
import React from 'react'
import {BsList} from 'react-icons/all'

export const List = () => {
    return (
        <>
            <IconContext.Provider value={{size: "2rem"}}>
                <BsList/>
            </IconContext.Provider>
        </>
    )
}