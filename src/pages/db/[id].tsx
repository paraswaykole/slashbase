import type { NextPage } from 'next'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import AppLayout from '../../components/layouts/applayout'

import DefaultErrorPage from 'next/error'
import { getDBConnection, getDBDataModels } from '../../redux/dbConnectionSlice'
import { useAppDispatch } from '../../redux/hooks'
import DBHomeFragment from '../../components/dbfragments/home'
import DBShowDataFragment from '../../components/dbfragments/showdata'

enum DBFragmentViewType {
    HOME = "HOME",
    SHOWDATA = "SHOWDATA"
}

const DBPage: NextPage = () => {

    const router = useRouter()
    const { id, mschema, mname } = router.query

    const [error404, setError404] = useState(false)
    const dispatch = useAppDispatch()

    const [dbFragmentViewType, setDBFragmentViewType] = useState<DBFragmentViewType>(DBFragmentViewType.HOME)

    useEffect(()=>{
        (async () => {
            if (id){
                try {
                    await dispatch((getDBConnection({dbConnId: String(id)}))).unwrap() 
                } catch (e){
                    console.log(e)
                    setError404(true)
                    return
                }
                dispatch((getDBDataModels({dbConnId: String(id)}))) 
            }
        })()
    }, [dispatch, router])

    useEffect(()=>{
        if(mschema && mname){
            setDBFragmentViewType(DBFragmentViewType.SHOWDATA) 
        } else {
            setDBFragmentViewType(DBFragmentViewType.HOME)
        }
    },[router])

    if(error404) {
        return (<DefaultErrorPage statusCode={404} />)
    }

    return (
        <AppLayout title="Home">
        <main className="maincontainer">
            {dbFragmentViewType === DBFragmentViewType.HOME && 
                <DBHomeFragment /> }
            {dbFragmentViewType === DBFragmentViewType.SHOWDATA && 
                <DBShowDataFragment /> }
        </main>
        </AppLayout>
    )
}

export default DBPage
