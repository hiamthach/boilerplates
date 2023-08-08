'use client';

import React from 'react';

import { Button, ButtonProps } from '@mantine/core';

const Btn = (props: ButtonProps) => {
  return <Button {...props}>{props.children}</Button>;
};

export default Btn;
