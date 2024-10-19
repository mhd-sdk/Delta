import { GaugeChart, GaugeTypes } from '@carbon/charts-react';
import '@carbon/charts-react/styles.css';
import { Slider } from '@carbon/react';
import { css } from '@emotion/css';
import { useState } from 'react';
import { useMarketMock } from './hooks/useMarketMock';

export const App = () => {
  const [buyAvg, setBuyAvg] = useState(50);
  const [sellAvg, setSellAvg] = useState(50);

  const { buyVolume, sellVolume } = useMarketMock({
    averageBuyVolume: buyAvg,
    averageSellVolume: sellAvg,
    buyFrequency: 50,
    sellFrequency: 50,
    duration: 300,
  });

  const totalVolume = buyVolume + sellVolume;
  const buyVolumePercentage = (buyVolume / totalVolume) * 100;

  return (
    <>
      <div className={styles.container}>
        <Slider
          labelText="Aggressive buyer (+-10)"
          value={buyAvg}
          min={0}
          max={100}
          stepMultiplier={50}
          onChange={(val) => setBuyAvg(val.value)}
          step={1}
          formatLabel={(val) => {
            if (val < 25) {
              return 'Low';
            } else if (val > 75) {
              return 'High';
            }
            return 'Medium';
          }}
        />
        <Slider
          labelText="Aggressive seller (+-10)"
          value={sellAvg}
          min={0}
          max={100}
          stepMultiplier={50}
          onChange={(val) => setSellAvg(val.value)}
          step={1}
          formatLabel={(val) => {
            if (val < 25) {
              return 'Low';
            } else if (val > 75) {
              return 'High';
            }
            return 'Medium';
          }}
        />
      </div>
      <GaugeChart
        data={[
          {
            group: 'value',
            value: buyVolumePercentage,
            min: -100,
            max: 100,
          },
          // {
          //   group: 'delta',
          //   value: -13.37,
          // },
        ]}
        options={{
          title: 'Delta',
          resizable: false,
          animations: true,
          height: '250px',
          gauge: {
            status: 'SUCCESS',
            numberFormatter: (value) => {
              switch (true) {
                case value > 75:
                  return 'High buy volume';
                case value < 25:
                  return 'High sell volume';
                case value > 50 && value < 75:
                  return 'Medium buy volume';
                case value > 25 && value < 50:
                  return 'Medium sell volume';
              }
              return '';
            },
            type: GaugeTypes.SEMI,
            showPercentageSymbol: false,
            alignment: 'center',
          },
        }}
      />
    </>
  );
};

const styles = {
  container: css`
    display: flex;
    flex-direction: column;
    gap: 16px;
    padding: 16px;
  `,
};
