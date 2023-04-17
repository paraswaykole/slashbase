import { useEffect } from "react"
import { Routes, Route, Link } from "react-router-dom"
import { Toaster } from 'react-hot-toast'
import Bowser from "bowser"
import { useAppDispatch, useAppSelector } from "./redux/hooks"
import { getProjects } from "./redux/projectsSlice"
import { getAllDBConnections } from "./redux/allDBConnectionsSlice"
import { getConfig } from "./redux/configSlice"
import AppLayout from "./components/layouts/applayout"
import HomePage from "./pages/home"
import ProjectPage from "./pages/project"
import NewDBPage from "./pages/project/newdb"
import DBPage from "./pages/db"
import AccountPage from "./pages/settings/account"
import AdvancedSettingsPage from "./pages/settings/advanced"
import AboutPage from "./pages/settings/about"
import SupportPage from "./pages/settings/support"
import GeneralSettingsPage from "./pages/settings/general"
import LogoutPage from "./pages/logout"
import { getUser, selectIsAuthenticated } from "./redux/currentUserSlice"


function App() {

  const isValidPlatform: boolean = Bowser.getParser(window.navigator.userAgent).getPlatformType(true) === "desktop"

  const dispatch = useAppDispatch()

  const isAuthenticated = useAppSelector(selectIsAuthenticated)

  useEffect(() => {
    if (isAuthenticated) {
      return
    }
    dispatch(getUser())
  }, [isAuthenticated, dispatch])

  useEffect(() => {
    if (isAuthenticated) {
      dispatch(getProjects())
      dispatch(getAllDBConnections({}))
      dispatch(getConfig())
    }
  }, [dispatch, isAuthenticated])

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
          <Route path="settings/account" element={<AccountPage />} />
          <Route path="settings/general" element={<GeneralSettingsPage />} />
          <Route path="settings/advanced" element={<AdvancedSettingsPage />} />
          <Route path="settings/about" element={<AboutPage />} />
          <Route path="settings/support" element={<SupportPage />} />
          <Route path="*" element={<NoMatch />} />
        </Route>
        <Route path="/logout" element={<LogoutPage />} />
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
