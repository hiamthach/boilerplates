'use client';

import React from 'react';

import { Button, ButtonProps } from 'antd';

const Btn = (props: ButtonProps) => {
  return <Button {...props}>{props.children}</Button>;
};

export default Btn;
