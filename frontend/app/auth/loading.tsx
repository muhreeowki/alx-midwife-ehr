"use client";

import React from "react";
import { TailSpin } from "react-loader-spinner";

const Loading = () => {
  return (
    <main className="flex items-center justify-center h-screen">
      <TailSpin
        visible={true}
        height="80"
        width="80"
        color="#CC0033"
        ariaLabel="tail-spin-loading"
        radius="3"
        wrapperStyle={{}}
        wrapperClass=""
      />
      <span className="sr-only">Loading...</span>
    </main>
  );
};

export default Loading;
