import { useEffect } from "react"
import { Routes, Route, Link } from "react-router-dom"
import { Toaster } from 'react-hot-toast'
import Bowser from "bowser"
import { useAppDispatch, useAppSelector } from "./redux/hooks"
import { connectLocal, selectIsConnected } from "./redux/apiSlice"
import { getProjects } from "./redux/projectsSlice"
import { getAllDBConnections } from "./redux/allDBConnectionsSlice"
import { getConfig } from "./redux/configSlice"
import AppLayout from "./components/layouts/applayout"
import HomePage from "./pages/home"
import ProjectPage from "./pages/project"
import NewDBPage from "./pages/project/newdb"
import DBPage from "./pages/db"
import DBHistoryPage from "./pages/db/history"
import DBPathPage from "./pages/db/path"
import DBQueryPage from "./pages/db/query"
import AdvancedSettingsPage from "./pages/settings/advanced"
import AboutPage from "./pages/settings/about"
import SupportPage from "./pages/settings/support"
import GeneralSettingsPage from "./pages/settings/general"


function App() {

  const isValidPlatform: boolean = Bowser.getParser(window.navigator.userAgent).getPlatformType(true) === "desktop"

  const dispatch = useAppDispatch()
  const isConnected = useAppSelector(selectIsConnected)

  useEffect(() => {
    const checkConnection = async () => {
      dispatch(connectLocal())
    }
    checkConnection()
  }, [dispatch])

  useEffect(() => {
    (async () => {
      if (isConnected) {
        await dispatch(getProjects())
        await dispatch(getAllDBConnections({}))
      }
      dispatch(getConfig())
    })()
  }, [dispatch, isConnected])

  if (!isValidPlatform) {
    return <NotSupportedPlatform />
  }

  return (
    <div className="appcontainer">
      <Routes>
        <Route path="/" element={<AppLayout />}>
          <Route index element={<HomePage />} />
          <Route path="project/:id" element={<ProjectPage />} />
          <Route path="project/:id/newdb" element={<NewDBPage />} />
          <Route path="db/:id" element={<DBPage />} />
          <Route path="db/:id/history" element={<DBHistoryPage />} />
          <Route path="db/:id/query/:queryId" element={<DBQueryPage />} />
          <Route path="db/:id/:path" element={<DBPathPage />} />
          <Route path="settings/general" element={<GeneralSettingsPage />} />
          <Route path="settings/advanced" element={<AdvancedSettingsPage />} />
          <Route path="settings/about" element={<AboutPage />} />
          <Route path="settings/support" element={<SupportPage />} />
          <Route path="*" element={<NoMatch />} />
        </Route>
      </Routes>
      <Toaster />
    </div>
  );
}

export default App

function NoMatch() {
  return (
    <div>
      <h2>Nothing to see here!</h2>
      <p>
        <Link to="/"><i className={`fas fa-home`} /> Go back to home</Link>
      </p>
    </div>
  );
}


function NotSupportedPlatform() {
  return (
    <div className="appcontainer">
      <h2>Slashbase is desktop only application!</h2>
      <p>Please use a desktop or laptop to continue...</p>
    </div>
  );
}
