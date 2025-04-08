import { createBrowserRouter, RouterProvider } from "react-router-dom"
import { Layouts } from "./Layout/Layout"
import { ROUTER_PATH } from "../shared/routes"
import { MainPage } from "../pages/MainPage"
import { RegisterPage } from "../pages/Auth/RegisterPage"
import { AuthPage } from "../pages/Auth/AuthPage/AuthPage"
import { CartPage } from "../pages/CartPage"

export function App() {
  const router = createBrowserRouter([
    {
      path: '/',
      element: <Layouts/>,
      children: [
        {
          path: ROUTER_PATH.MAIN,
          element: <MainPage/>
        },
        {
          path: ROUTER_PATH.REGISTER,
          element: <RegisterPage/>
        },
        {
          path: ROUTER_PATH.AUTH,
          element: <AuthPage/>
        },
        {
          path: ROUTER_PATH.CART,
          element: <CartPage/>
        },
      ]
  }
])

  return (
    <RouterProvider router={router}/>
  )
}

