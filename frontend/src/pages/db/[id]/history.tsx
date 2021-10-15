import type { NextPage } from 'next'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import AppLayout from '../../../components/layouts/applayout'

import DefaultErrorPage from 'next/error'
import { getDBConnection, getDBDataModels, getDBQueries, selectDBConnection } from '../../../redux/dbConnectionSlice'
import { useAppDispatch, useAppSelector } from '../../../redux/hooks'
import { DBConnection } from '../../../data/models'
import DBHistoryFragment from '../../../components/dbfragments/history'

const DBHistoryPage: NextPage = () => {

    const router = useRouter()
    const { id } = router.query
    
    const dispatch = useAppDispatch()
    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)

    const [error404, setError404] = useState(false)

    useEffect(()=>{
        (async () => {
            if (id){
                try {
                    await dispatch((getDBConnection({dbConnId: String(id)}))).unwrap() 
                } catch (e){
                    setError404(true)
                    return
                }
                dispatch((getDBDataModels({dbConnId: String(id)}))) 
                dispatch((getDBQueries({dbConnId: String(id)})))
            }
        })()
    }, [dispatch, router])


    if(error404) {
        return (<DefaultErrorPage statusCode={404} />)
    }

    return (
        <AppLayout title={dbConnection ? dbConnection.name + " | Slashbase" : "Slashbase"}>
           <DBHistoryFragment />
        </AppLayout>
    )
}

export default DBHistoryPage
