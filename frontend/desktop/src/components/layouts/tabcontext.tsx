import { createContext } from 'react';
import { Tab } from '../../data/models'

const TabContext = createContext<Tab | undefined>(undefined);

export default TabContext