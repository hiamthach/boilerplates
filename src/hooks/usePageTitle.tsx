import useIsomorphicLayoutEffect from './useIsomorphicLayoutEffect';

import { APP_NAME } from '@/config/constants/env.const';

function usePageTitle(title: string): void {
  useIsomorphicLayoutEffect(() => {
    window.document.title = title + ' | ' + APP_NAME;
  }, [title]);
}

export default usePageTitle;
