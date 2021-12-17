import { ColorModeScript } from "@chakra-ui/react"
import * as React from "react"
import ReactDOM from "react-dom"
import {BrowserRouter, Routes, Route} from "react-router-dom"
import { App as A } from "./App"
import reportWebVitals from "./reportWebVitals"
import * as serviceWorker from "./serviceWorker"

ReactDOM.render(
  <React.StrictMode>
    <ColorModeScript />
    <App />
  </React.StrictMode>
  , document.getElementById("root"),
)

function App () {
  return <BrowserRouter>
    <Routes>
      <Route path="/" element={<A />}/>
      <Route path="/auth" element={<Auth />}/>
      <Route path="*" element={<NotFound />}/>
    </Routes>
  </BrowserRouter>
}

function Auth () {
  return <h1>Authorization page</h1>
}

function NotFound () {
  return <h1>Page not Found</h1>
}

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://cra.link/PWA
serviceWorker.unregister()

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
