import type { NextPage } from 'next'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import DBConnCard from '../../components/dbconncard/dbconncard'
import AppLayout from '../../components/layouts/applayout'
import { DBConnection } from '../../data/models'
import apiService from '../../network/apiService'

const ProjectPage: NextPage = () => {

  const router = useRouter()

  const [databases, setDatabases] = useState<DBConnection[]>([])

  useEffect(()=>{
    (async () => {
      const { id } = router.query    
      let response = await apiService.getDBConnectionsByProject(String(id))
      if(response.success){
        setDatabases(response.data)
      }
    })()
  }, [])

  return (
    <AppLayout title="Home">
      <main className="maincontainer">
        <h1>All Databases</h1>
        {databases.map((db: DBConnection) => (
          <DBConnCard key={db.id} dbConn={db}/>
        ))}
      </main>
    </AppLayout>
  )
}

export default ProjectPage
