import React, { useState, useEffect } from 'react'
import { signIn, csrfToken, useSession } from "next-auth/react"
import Router from 'next/router'
import * as Yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { useForm } from 'react-hook-form';
import Swal from 'sweetalert2'

const signin = ({ csrfToken }) => {
  const { data: session } = useSession()

 

  const validatorSchema = Yup.object().shape({
    username: Yup.string().required('Username is required'),
    password: Yup.string().required('Password is required')
  })
  const formOptions = { resolver: yupResolver(validatorSchema) }
  const { register, handleSubmit, formState } = useForm(formOptions)
  const { errors } = formState;


  var count = Object.keys(errors).length
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: (toast) => {
      toast.addEventListener('mouseenter', Swal.stopTimer)
      toast.addEventListener('mouseleave', Swal.resumeTimer)
    }
  })


  if (count > 0) {
    const errorReverse = [errors].reverse()
    errorReverse.map((error, index) => {
      Toast.fire({
        icon: 'error',
        title: error[Object.keys(error)[0]].message
      })
    }
    )
  }

  const onSubmit = async (e) => {
    let options = { redirect: false, username: e.username, password: e.password, callbackUrl: `${window.location.origin}/` }
    const res = await signIn('username-login', options)
    if (res.ok && res.status === 200) {
      Toast.fire({
        icon: 'success',
        title: 'Sign in success'
      }).then(() => {
        Router.push('/')
      })
    } else {
      Toast.fire({
        icon: 'error',
        title: 'Sign in failed, username or password is incorrect'
      })
    }
  }

  return (

    <div className="bg-white dark:bg-gray-900">
      <div className="flex justify-center h-screen">
        <div
          className="hidden bg-cover lg:block lg:w-2/3"
          style={{
            backgroundImage:
              "url(https://images.unsplash.com/photo-1616763355603-9755a640a287?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1470&q=80)"
          }}
        >
          <div className="flex items-center h-full px-20 bg-gray-900 bg-opacity-40">
            <div>
              <h2 className="text-4xl font-bold text-white">Brand</h2>
              <p className="max-w-xl mt-3 text-gray-300">
                {
                  session ? "<div>You are already signed in</div>" : "<div>Sign in to your account</div>"
                }

              </p>
            </div>
          </div>
        </div>
        <div className="flex items-center w-full max-w-md px-6 mx-auto lg:w-2/6">
          <div className="flex-1">
            <div className="text-center">
              <h2 className="text-4xl font-bold text-center text-gray-700 dark:text-white">
                Service System Management
              </h2>
              <p className="mt-3 text-gray-500 dark:text-gray-300">
                Sign in to access your account
              </p>
            </div>
            <div className="mt-8">
              <form onSubmit={handleSubmit(onSubmit)}>
                <div>
                  <label
                    htmlFor="email"
                    className="block mb-2 text-sm text-gray-600 dark:text-gray-200"
                  >
                    Username
                  </label>
                  <input
                    {...register('username')}
                    type="text"
                    name="username"
                    id="username"
                    placeholder="your username"
                    className="block w-full px-4 py-2 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-md dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40"
                  />
                </div>
                <div className="mt-6">
                  <div className="flex justify-between mb-2">
                    <label
                      htmlFor="password"

                      className="text-sm text-gray-600 dark:text-gray-200"
                    >
                      Password
                    </label>
                    <a
                      href="#"
                      className="text-sm text-gray-400 focus:text-blue-500 hover:text-blue-500 hover:underline"
                    >
                      Forgot password?
                    </a>
                  </div>
                  <input
                    type="password"
                    name="password"
                    id="password"
                    {...register('password')}
                    placeholder="your password"
                    className="block w-full px-4 py-2 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-md dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40"
                  />
                </div>
                <div className="mt-6">
                  <button
                    className="w-full px-4 py-2 tracking-wide text-white transition-colors duration-200 transform bg-blue-500 rounded-md hover:bg-blue-400 focus:outline-none focus:bg-blue-400 focus:ring focus:ring-blue-300 focus:ring-opacity-50"
                    type='submit'>
                    Sign in
                  </button>
                </div>
              </form>
              <p className="mt-6 text-sm text-center text-gray-400">
                Don't have an account yet?{" "}
                <a
                  href="#"
                  onClick={() => Router.push("/auth/signup")}
                  className="text-blue-500 focus:outline-none focus:underline hover:underline"
                >
                  Sign up
                </a>
                .
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>

  )
}

export async function getInitalProps(context) {
  return {
    csrfToken: await csrfToken(context)
  }
}

export default signin