import type { Meta, StoryObj } from '@storybook/react';

import Btn from './Btn';

const meta = {
  title: 'Shared/Btn',
  component: Btn,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof Btn>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    children: 'Default',
    variant: 'filled',
  },
};

export const Outline: Story = {
  args: {
    children: 'Outline',
    variant: 'outline',
  },
};

export const Large: Story = {
  args: {
    children: 'Large',
    size: 'lg',
  },
};

export const Small: Story = {
  args: {
    children: 'Small',
    size: 'sm',
  },
};
