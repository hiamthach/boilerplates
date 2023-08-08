import type { Meta, StoryObj } from '@storybook/react';

import Btn from './Btn';

const meta: Meta<typeof Btn> = {
  component: Btn,
};

export default meta;
type Story = StoryObj<typeof Btn>;

export const FirstStory: Story = {
  args: {
    children: 'Button',
  },
};
