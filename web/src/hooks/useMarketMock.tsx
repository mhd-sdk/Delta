import { useEffect, useState } from 'react';

export interface MarketMockParams {
  averageBuyVolume: number; // Average volume for buy orders
  buyFrequency: number; // Frequency of buy orders (milliseconds)
  averageSellVolume: number; // Average volume for sell orders
  sellFrequency: number; // Frequency of sell orders (milliseconds)
  duration: number; // x last seconds to analyze
}

export interface Order {
  side: OrderSide;
  size: number;
  time: Date;
  price: number;
}
export type Tape = Order[];

export enum OrderSide {
  Buy = 'buy',
  Sell = 'sell',
}

function generateRandomVolume(mean: number, deviation: number): number {
  const randomOffset = Math.random() * deviation * 2 - deviation;
  return mean + randomOffset;
}

export const useMarketMock = (
  params: MarketMockParams
): {
  delta: number;
  tape: Tape;
  buyVolume: number;
  sellVolume: number;
} => {
  const [tape, setTape] = useState<Tape>([]);
  const [delta, setDelta] = useState<number>(0);
  const [totalBuyVolume, setTotalBuyVolume] = useState<number>(0);
  const [totalSellVolume, setTotalSellVolume] = useState<number>(0);

  const generateRandomOrder = (side: OrderSide, averageSize: number) => {
    const size = generateRandomVolume(averageSize, 2);
    const price = Math.random() * 100;
    const order: Order = {
      side,
      size,
      time: new Date(),
      price,
    };
    return order;
  };

  useEffect(() => {
    // Mock buy orders
    const buyInterval = setInterval(() => {
      const buyOrder = generateRandomOrder(
        OrderSide.Buy,
        params.averageBuyVolume
      );
      setTape((prevTape) => [...prevTape, buyOrder]);
    }, params.buyFrequency);

    // Mock sell orders
    const sellInterval = setInterval(() => {
      const sellOrder = generateRandomOrder(
        OrderSide.Sell,
        params.averageSellVolume
      );
      setTape((prevTape) => [...prevTape, sellOrder]);
    }, params.sellFrequency);

    return () => {
      clearInterval(buyInterval);
      clearInterval(sellInterval);
    };
  }, [params]);

  useEffect(() => {
    // return delta of last x seconds
    const lastStudiedTime = new Date();
    lastStudiedTime.setMilliseconds(
      lastStudiedTime.getMilliseconds() - params.duration
    );

    const lastOrders = tape.filter((order) => order.time >= lastStudiedTime);

    const buyVolumes = lastOrders.filter(
      (order) => order.side === OrderSide.Buy
    );
    const totalBuyVolume = buyVolumes.reduce(
      (acc, order) => acc + order.size,
      0
    );

    const sellVolumes = lastOrders.filter(
      (order) => order.side === OrderSide.Sell
    );
    const totalSellVolume = sellVolumes.reduce(
      (acc, order) => acc + order.size,
      0
    );

    setDelta(totalBuyVolume - totalSellVolume);
    setTotalBuyVolume(totalBuyVolume);
    setTotalSellVolume(totalSellVolume);
  }, [tape, params]);

  return {
    tape,
    delta,
    buyVolume: totalBuyVolume,
    sellVolume: totalSellVolume,
  };
};
