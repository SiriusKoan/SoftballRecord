import './App.css';
import styled from '@emotion/styled'
import { useState, useEffect, useCallback } from 'react'

const GuestTeam = "Yankees"
const HomeTeam = "Red Sox"

const OutsContainerStyle = styled.div`
  border-radius: 50%;
  display: inline-block;
  margin-up: 10px;
  margin-bottom: 10px;
  align-items: center;
  text-align: center;
  width: 100%;
`

const PaFormContainerStyle = styled.div`
  border-radius: 50%;
  margin-up: 10px;
  margin-bottom: 10px;
  align-items: center;
  text-align: left;
  width: 50%;
  float:left;
`

const HistoryContainerStyle = styled.div`
  border-radius: 50%;
  margin-up: 10px;
  margin-bottom: 10px;
  align-items: right;
  text-align: right;
  width: 50%;
  float:left;
`

const InputStyle = styled.div`
  font-font-size-adjust: 0.90;
  font-family: 微軟正黑體;
  margin-left: 20px;
  margin-right: 20px;
`
const HeaderStyle = styled.div`
  font-color: #010101;
  font-size: 40px;
  width: 100%;
  text-align: center;
  ;
`
const ScoreboardStyle = styled.div`
 display: flex;
 background-color: pink;
 width: 100%;
 text-align: center;
 border:1px solid #000;
 ;
`
const TableStyle = styled.div`
  font-size: 30px;
  border-collapse: collapse;
  font-family: 微軟正黑體;
  align-content: center;
  align-self: center;
  align-items: center;
  text-align: center;
`
const HistoryTableStyle = styled.div`
  font-size: 15px;
  border-collapse: collapse;
  font-family: 微軟正黑體;
  align-content: center;
  align-self: center;
  align-items: center;
  text-align: center;
`
const ButtonStyle = styled.button`
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
  border-radius: 10px;
  color: #012311;
  height: 40px;
  width: 100px;
  margin-left: 20px;
  margin-right: 20px;
  margin-up: 40px;
  margin-bottom: 40px;
  font-size: 20px;
  text-align: center;
  background-color: ${({bgcolor}) => bgcolor};
  &:hover {
    background-color: lightblue;
  };
  
`
const CurrentInningStyle = styled.div`
  color: red;
`
const NotCurrentInningStyle = styled.div`
  color: black;
`
const OutSign = styled.div`
  display: inline-block;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  margin-right: 5px;
  background-color: lightgrey;
`

const OutSign2 = styled.div`
  display: inline-block;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  margin-right: 5px;
  background-color: yellow;
`

const Header = () => (
  <h1>
    <HeaderStyle>
      比賽紀錄
    </HeaderStyle>
  </h1>
)

const Inning = ({ currentInning, targetInning }) => {
  if (currentInning === targetInning) {
    return (
      <td width="100px">
        <CurrentInningStyle>
          {targetInning}
        </CurrentInningStyle>
      </td>
    )
  } else {
    return (
      <td width="100px">
        <NotCurrentInningStyle>
          {targetInning}
        </NotCurrentInningStyle>
      </td>
    )
  }
}

const PrintGuestTeam = ({ isBottom }) => {
  if (isBottom === false) {
    return (
        <td><font color="red">{GuestTeam}</font></td>
    )
  } else {
    return (
        <td key={GuestTeam}>{GuestTeam}</td>
    )
  }
}

const PrintHomeTeam = ({ isBottom }) => {
  if (isBottom === true) {
    return (
      <td><font color="red">{HomeTeam}</font></td>
    )
  } else {
    return (  
        <td key={HomeTeam}>{HomeTeam}</td>
    )
  }
}

const InningsRow = ({ currentInning, GuestScores }) => {
  const Innings = Array.from({ length: 9 }, (_, index) => index);
  return (
    <tr>
      <td></td>
      {Innings.map((score, index) => (
        <Inning key={index} currentInning={currentInning} targetInning={index + 1} />
      ))}

      <td width="70px">R</td>
      <td width="70px">H</td>
      <td width="70px">E</td>
    </tr>
  )
}

const Score = ({ score }) => {
  console.log("score: ", score)
  if (score === -1) {
    return (
      <td></td>
    )
  } else {
    return (
      <td>
        {score}
      </td>
    )
  }
}

const ScoreBoard = (props) => {
  const { currentInning, isBottom, scores } = props
  // const Innings = Array.from({ length: 9 }, (_, index) => index);
  console.log("scores[0]: ", scores[0])

  return (
    <div>
      <TableStyle>
        <table border="1">
          <tbody>
            <InningsRow currentInning={currentInning} />
            <tr>
              <PrintGuestTeam isBottom={isBottom} GuestScores={scores[0]} />
              {scores[0].map((score, index) => (
                <Score key={index} score={score} />
              ))}
            </tr>
            <tr>
              <PrintHomeTeam isBottom={isBottom} />
              {scores[1].map((score, index) => (
                <Score key={index} score={score} />
              ))}
            </tr>
          </tbody>
        </table>
      </TableStyle>
    </div>
  )
}

const OutBoard = ({ currentOuts }) => {
  const O1 = Array.from({length: currentOuts}, (_, index) => index)
  const O2 = Array.from({length: 3 - currentOuts}, (_, index) => index)
  return (
    <OutsContainerStyle>
      <div>
        <h3 align="center">
          Outs
        </h3>
        {O1.map((op, index) => (
          <OutSign2 key={index}> </OutSign2>
        ))}
        {O2.map((op, index) => (
          <OutSign key={index+3}></OutSign>
        ))}
      </div>
    </OutsContainerStyle>
  )
}

const ButtonArea = ({ changeInning }) => {

  return (
    <>
      <ButtonStyle onClick={changeInning} bgcolor={'#66B3FF'}>
        Change
      </ButtonStyle>
      <ButtonStyle bgcolor={'#02DF82'}>
        Export
      </ButtonStyle>
      <ButtonStyle bgcolor={'#FF8040'}>
        Reset
      </ButtonStyle>
    </>
  )
}

const BattingOrder = ({ handleChange0, setBattingOrder }) => {
  const Options = Array.from({ length: 11 }, (_, index) => index + 1)
  
 const handleChange = event => setBattingOrder(event.target.value)
  
  return (
    <InputStyle>
      <label>棒次: </label>
      <select onChange={handleChange}>
        <option></option>
        {Options.map((op, index) => (
          <option key={index}>{op}</option>
        ))}
      </select>
    </InputStyle>
  )
}

const PlayerNumber = ({ setPlayerNumber }) => {

  const handleChange = event => setPlayerNumber(event.target.value)

  return (
    <InputStyle>
      <label>背號: </label>
      <input onChange={handleChange} />
    </InputStyle>
  )
}

const Direction = ({ setDirection }) => {
  const PositionCode = ['', '1 (投手)', '2 (捕手)', '3 (一壘)', '4 (二壘)', '5 (三壘)', '6 (游擊)', '7 (左外)', '8 (中外)', '9 (右外)', '10 (自由)'];
  
  const handleChange = event => setDirection(event.target.value)

  return (
    <InputStyle>
      <label>打擊方向: </label>
      <select onChange={handleChange}>
        {PositionCode.map((position, index) => (
          <option key={index}>
            {position}
          </option>
        ))}
      </select>
    </InputStyle>
  )
}

const BattingResult = ({ setBattingResult }) => {
  const Result = ['', 'GO (滾地出局)', 'FO (飛球出局)', 'FC (野手選擇)', 'E (對手失誤)', 'K (被三振)', 'BB (保送)', 'SF (高飛犧牲打)', '1B (一壘安打)', '2B (二壘安打)', '3B (三壘安打)', 'HR (全壘打)'];

  const handleChange = event => setBattingResult(event.target.value)

  return (
    <InputStyle>
      <label>打擊結果: </label>
      <select onChange={handleChange}>
        {Result.map((result, index) => (
          <option key={index}>
            {result}
          </option>
        ))}
      </select>
    </InputStyle>
  )
}

const Outs = ({ setOuts }) => {
  const Result = ['', '0', '1', '2', '3'];

  const handleChange = event => setOuts(event.target.value)

  return (
    <InputStyle>
      <label>出局: </label>
      <select onChange={handleChange}>
        {Result.map((result, index) => (
          <option key={index}>
            {result}
          </option>
        ))}
      </select>
    </InputStyle>
  )
}

const Rbi = ({ setRbi }) => {
  const Result = [0, 1, 2, 3, 4];

  const handleChange = event => setRbi(parseInt(event.target.value))

  return (
    <InputStyle>
      <label>打點: </label>
      <select onChange={handleChange}>
        {Result.map((result, index) => (
          <option key={index}>
            {result}
          </option>
        ))}
      </select>
    </InputStyle>
  )
}
/*
const PaFormArea = ({ setHistory, history,
                      setCurrentPa, currentPa, 
                      setBattingOrder, battingOrder, 
                      setPlayerNumber, playerNumber,
                      setDirection, direction,
                      setBattingResult, battingResult,
                      setOuts, outs, 
                      setRbi, rbi,
                      setCurrentOuts, currentOuts, isBottom}) => {


  const handleClick = () => {
    setCurrentOuts(outs)
    setCurrentPa(currentPa+1)
    setHistory([...history, {
      team: isBottom ? HomeTeam : GuestTeam,
      battingOrder: battingOrder,
      playerNumber: playerNumber,
      direction: direction,
      battingResult: battingResult,
      outs: outs,
      rbi: rbi,
      currentOuts: currentOuts
    }])
  }

  useEffect(() => {
    console.log(history)
  }, [history])

  return (
    <PaFormContainerStyle>
      <ButtonStyle onClick={handleClick}> 送出 </ButtonStyle>
      <font>打席數: {currentPa}</font>
      <BattingOrder setBattingOrder={setBattingOrder} />
      <PlayerNumber setPlayerNumber={setPlayerNumber} />
      <Direction setDirection={setDirection} />
      <BattingResult setBattingResult={setBattingResult} />
      <Outs setOuts={setOuts} />
      <Rbi setRbi={setRbi} />
    </PaFormContainerStyle>
  )
}
*/

const HistoryArea = ({ history }) => {
  console.log(history)
  return (
    <HistoryContainerStyle>
      <HistoryTableStyle>
      <table>
      <thead>History
      </thead>
      <tbody>
      <tr>
        <td>隊伍</td>
        <td>局數</td>
        <td>棒次</td>
        <td>背號</td>
        <td>打擊方向</td>
        <td>打擊結果</td>
        <td>出局數</td>
        <td>打點</td>
      </tr>
      {history.map((hist, index) => (
              <tr>
                <td>{hist.team}</td>
                <td>{hist.inning}</td>
                <td>{hist.battingOrder}</td>
                <td>{hist.playerNumber}</td>
                <td>{hist.direction}</td>
                <td>{hist.battingResult}</td>
                <td>{hist.outs}</td>
                <td>{hist.rbi} </td>
              </tr>
            
          
        ))}
        </tbody>
        </table>
        </HistoryTableStyle>
    </HistoryContainerStyle>
  )
}

function RecordApp() {
  const [history, setHistory] = useState([])
  const [currentInning, setCurrentInning] = useState(1)
  const [currentOuts, setCurrentOuts] = useState(0)
  const [currentPa, setCurrentPa] = useState(0)
  const [battingOrder, setBattingOrder] = useState(1)
  const [playerNumber, setPlayerNumber] = useState("")
  const [direction, setDirection] = useState("")
  const [battingResult, setBattingResult] = useState("")
  const [outs, setOuts] = useState("")
  const [rbi, setRbi] = useState(0)
  const [accuScore, setAccuScore] = useState(0)

  const [isBottom, setBottom] = useState(0);
  const [scores, setScores] = useState(
    [
      Array.from({ length: 9 }, (_) => -1), 
      Array.from({ length: 9 }, (_) => -1)
    ]
    )
  const changeInning = () => {
    scores[isBottom][currentInning-1] = accuScore
    setAccuScore(0)
    setScores(scores)
    if (isBottom === 0) {
      setBottom(1)
    }
    else if (currentInning === 9) {
      setBottom(0)
      setCurrentInning(1);
    } else {
      setBottom(0)
      setCurrentInning(currentInning + 1);
    }
  }

  const setByLastHistory = () => {
    setCurrentPa(history[history.length-1].currentPa)
    setCurrentOuts(history[history.length-1].outs)
  }

  const handleSendClick = () => {

    setHistory([...history, {
      team: isBottom ? HomeTeam : GuestTeam,
      inning: currentInning,
      battingOrder: battingOrder,
      playerNumber: playerNumber,
      direction: direction,
      battingResult: battingResult,
      outs: outs,
      rbi: rbi,
      currentOuts: currentOuts
    }])

  }

  const handleUndoClick = () => {
    if (history.length === 0) {
      return
    }
    setHistory(history.filter((_, i) => {return i != history.length-1}))
    setByLastHistory();
  }

  

  useEffect(() => {
    if (history.length === 0) {
      setCurrentPa(0)
      setCurrentOuts(0)
      return
    }
    setCurrentPa(history[history.length-1].currentPa)
    setCurrentOuts(history[history.length-1].outs)
    setAccuScore(accuScore + history[history.length-1].rbi)
  }, [history])

  

  console.log(battingOrder);
  return (
    <>
      <div>
        <Header />
        <ButtonArea changeInning={changeInning} />
        <ScoreboardStyle />
        <ScoreBoard
          currentInning={currentInning}
          isBottom={isBottom}
          scores={scores}
          />
      </div>
      <OutsContainerStyle>
        <OutBoard currentOuts={currentOuts} />
      </OutsContainerStyle>

      
      <PaFormContainerStyle>
        <ButtonStyle onClick={handleSendClick}> 送出 </ButtonStyle>
        <ButtonStyle onClick={handleUndoClick}> Undo </ButtonStyle>
        <p>打席數: {currentPa}</p>
        <BattingOrder setBattingOrder={setBattingOrder} />
        <PlayerNumber setPlayerNumber={setPlayerNumber} />
        <Direction setDirection={setDirection} />
        <BattingResult setBattingResult={setBattingResult} />
        <Outs setOuts={setOuts} />
        <Rbi setRbi={setRbi} />
      </PaFormContainerStyle>
      
      
      <HistoryArea history={history} isBottom={isBottom} />
      
    </>
  );
}

export default RecordApp;