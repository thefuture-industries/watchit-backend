import { useEffect, useState } from 'react';

export default function useServerData(fn: () => any) {
  const [data, setData] = useState<any>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const sendData = async () => {
    setIsLoading(true);

    try {
      const response = await fn();

      setData(response);
    } catch (error: any) {
      throw new Error(error);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    sendData();
  }, []);

  const refetch = () => sendData();

  return { data, isLoading, refetch };
}
