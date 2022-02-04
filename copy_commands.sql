COPY "public"."BCNUser" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUser.csv' DELIMITER ',';

COPY "public"."BCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNAccount.csv' DELIMITER ',';

COPY "public"."DefaultBCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/DefaultBCNAccount.csv' DELIMITER ',';

COPY "public"."Password" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Password.csv' DELIMITER ',';

COPY "public"."LeoLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LeoLLT.csv' DELIMITER ',';

COPY "public"."BCNUserLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUserLLT.csv' DELIMITER ',';

COPY "public"."Transaction"("id","amount","senderBcnAccountId","recipientBcnAccountId","resultedAt","exchangeRateMultiplier","exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Transaction.csv' DELIMITER ',' NULL 'null';


COPY "public"."SendMoneyWithinBCN"("id", "recipientBcnAccountId", "amount", "senderBcnAccountId","claimedFeeCustomer", "claimedFeeExpiresAt", "status","transactionId", "feeCustomerTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyWithinBCN.csv' DELIMITER ',';

COPY "public"."SendMoneyAccountToAccount"("id", "amount", "senderBcnAccountId", "recipientBcnAccountId", "claimedFeeCustomer","confirmationExpiresAt", "status", "transactionId","feeCustomerTransactionId", "exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyAccountToAccount.csv' DELIMITER ',';

COPY "public"."CounterpartyTransaction"("id", "counterpartyFinalTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/CounterPartyTransaction.csv' DELIMITER ',';

COPY "public"."LoadMoneyMPGS"("id", "amount", "recipientAccountId", "counterpartyId", "claimedFeeCounterparty","claimedFeeCustomer", "claimedFeeExpiresAt", "status", "counterpartyTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LoadMoneyMPGS.csv' DELIMITER ',';
COPY "public"."BCNUser" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUser.csv' DELIMITER ',';

COPY "public"."BCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNAccount.csv' DELIMITER ',';

COPY "public"."DefaultBCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/DefaultBCNAccount.csv' DELIMITER ',';

COPY "public"."Password" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Password.csv' DELIMITER ',';

COPY "public"."LeoLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LeoLLT.csv' DELIMITER ',';

COPY "public"."BCNUserLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUserLLT.csv' DELIMITER ',';

COPY "public"."Transaction"("id","amount","senderBcnAccountId","recipientBcnAccountId","resultedAt","exchangeRateMultiplier","exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Transaction.csv' DELIMITER ',' NULL 'null';


COPY "public"."SendMoneyWithinBCN"("id", "recipientBcnAccountId", "amount", "senderBcnAccountId","claimedFeeCustomer", "claimedFeeExpiresAt", "status","transactionId", "feeCustomerTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyWithinBCN.csv' DELIMITER ',';

COPY "public"."SendMoneyAccountToAccount"("id", "amount", "senderBcnAccountId", "recipientBcnAccountId", "claimedFeeCustomer","confirmationExpiresAt", "status", "transactionId","feeCustomerTransactionId", "exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyAccountToAccount.csv' DELIMITER ',';

COPY "public"."CounterpartyTransaction"("id", "counterpartyFinalTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/CounterPartyTransaction.csv' DELIMITER ',';

COPY "public"."LoadMoneyMPGS"("id", "amount", "recipientAccountId", "counterpartyId", "claimedFeeCounterparty","claimedFeeCustomer", "claimedFeeExpiresAt", "status", "counterpartyTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LoadMoneyMPGS.csv' DELIMITER ',';
COPY "public"."BCNUser" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUser.csv' DELIMITER ',';

COPY "public"."BCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNAccount.csv' DELIMITER ',';

COPY "public"."DefaultBCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/DefaultBCNAccount.csv' DELIMITER ',';

COPY "public"."Password" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Password.csv' DELIMITER ',';

COPY "public"."LeoLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LeoLLT.csv' DELIMITER ',';

COPY "public"."BCNUserLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUserLLT.csv' DELIMITER ',';

COPY "public"."Transaction"("id","amount","senderBcnAccountId","recipientBcnAccountId","resultedAt","exchangeRateMultiplier","exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Transaction.csv' DELIMITER ',' NULL 'null';


COPY "public"."SendMoneyWithinBCN"("id", "recipientBcnAccountId", "amount", "senderBcnAccountId","claimedFeeCustomer", "claimedFeeExpiresAt", "status","transactionId", "feeCustomerTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyWithinBCN.csv' DELIMITER ',';

COPY "public"."SendMoneyAccountToAccount"("id", "amount", "senderBcnAccountId", "recipientBcnAccountId", "claimedFeeCustomer","confirmationExpiresAt", "status", "transactionId","feeCustomerTransactionId", "exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyAccountToAccount.csv' DELIMITER ',';

COPY "public"."CounterpartyTransaction"("id", "counterpartyFinalTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/CounterPartyTransaction.csv' DELIMITER ',';

COPY "public"."LoadMoneyMPGS"("id", "amount", "recipientAccountId", "counterpartyId", "claimedFeeCounterparty","claimedFeeCustomer", "claimedFeeExpiresAt", "status", "counterpartyTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LoadMoneyMPGS.csv' DELIMITER ',';
COPY "public"."BCNUser" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUser.csv' DELIMITER ',';

COPY "public"."BCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNAccount.csv' DELIMITER ',';

COPY "public"."DefaultBCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/DefaultBCNAccount.csv' DELIMITER ',';

COPY "public"."Password" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Password.csv' DELIMITER ',';

COPY "public"."LeoLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LeoLLT.csv' DELIMITER ',';

COPY "public"."BCNUserLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUserLLT.csv' DELIMITER ',';

COPY "public"."Transaction"("id","amount","senderBcnAccountId","recipientBcnAccountId","resultedAt","exchangeRateMultiplier","exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Transaction.csv' DELIMITER ',' NULL 'null';


COPY "public"."SendMoneyWithinBCN"("id", "recipientBcnAccountId", "amount", "senderBcnAccountId","claimedFeeCustomer", "claimedFeeExpiresAt", "status","transactionId", "feeCustomerTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyWithinBCN.csv' DELIMITER ',';

COPY "public"."SendMoneyAccountToAccount"("id", "amount", "senderBcnAccountId", "recipientBcnAccountId", "claimedFeeCustomer","confirmationExpiresAt", "status", "transactionId","feeCustomerTransactionId", "exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyAccountToAccount.csv' DELIMITER ',';

COPY "public"."CounterpartyTransaction"("id", "counterpartyFinalTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/CounterPartyTransaction.csv' DELIMITER ',';

COPY "public"."LoadMoneyMPGS"("id", "amount", "recipientAccountId", "counterpartyId", "claimedFeeCounterparty","claimedFeeCustomer", "claimedFeeExpiresAt", "status", "counterpartyTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LoadMoneyMPGS.csv' DELIMITER ',';
COPY "public"."BCNUser" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUser.csv' DELIMITER ',';

COPY "public"."BCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNAccount.csv' DELIMITER ',';

COPY "public"."DefaultBCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/DefaultBCNAccount.csv' DELIMITER ',';

COPY "public"."Password" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Password.csv' DELIMITER ',';

COPY "public"."LeoLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LeoLLT.csv' DELIMITER ',';

COPY "public"."BCNUserLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUserLLT.csv' DELIMITER ',';

COPY "public"."Transaction"("id","amount","senderBcnAccountId","recipientBcnAccountId","resultedAt","exchangeRateMultiplier","exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Transaction.csv' DELIMITER ',' NULL 'null';


COPY "public"."SendMoneyWithinBCN"("id", "recipientBcnAccountId", "amount", "senderBcnAccountId","claimedFeeCustomer", "claimedFeeExpiresAt", "status","transactionId", "feeCustomerTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyWithinBCN.csv' DELIMITER ',';

COPY "public"."SendMoneyAccountToAccount"("id", "amount", "senderBcnAccountId", "recipientBcnAccountId", "claimedFeeCustomer","confirmationExpiresAt", "status", "transactionId","feeCustomerTransactionId", "exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyAccountToAccount.csv' DELIMITER ',';

COPY "public"."CounterpartyTransaction"("id", "counterpartyFinalTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/CounterPartyTransaction.csv' DELIMITER ',';

COPY "public"."LoadMoneyMPGS"("id", "amount", "recipientAccountId", "counterpartyId", "claimedFeeCounterparty","claimedFeeCustomer", "claimedFeeExpiresAt", "status", "counterpartyTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LoadMoneyMPGS.csv' DELIMITER ',';
COPY "public"."BCNUser" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUser.csv' DELIMITER ',';

COPY "public"."BCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNAccount.csv' DELIMITER ',';

COPY "public"."DefaultBCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/DefaultBCNAccount.csv' DELIMITER ',';

COPY "public"."Password" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Password.csv' DELIMITER ',';

COPY "public"."LeoLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LeoLLT.csv' DELIMITER ',';

COPY "public"."BCNUserLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUserLLT.csv' DELIMITER ',';

COPY "public"."Transaction"("id","amount","senderBcnAccountId","recipientBcnAccountId","resultedAt","exchangeRateMultiplier","exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Transaction.csv' DELIMITER ',' NULL 'null';


COPY "public"."SendMoneyWithinBCN"("id", "recipientBcnAccountId", "amount", "senderBcnAccountId","claimedFeeCustomer", "claimedFeeExpiresAt", "status","transactionId", "feeCustomerTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyWithinBCN.csv' DELIMITER ',';

COPY "public"."SendMoneyAccountToAccount"("id", "amount", "senderBcnAccountId", "recipientBcnAccountId", "claimedFeeCustomer","confirmationExpiresAt", "status", "transactionId","feeCustomerTransactionId", "exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyAccountToAccount.csv' DELIMITER ',';

COPY "public"."CounterpartyTransaction"("id", "counterpartyFinalTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/CounterPartyTransaction.csv' DELIMITER ',';

COPY "public"."LoadMoneyMPGS"("id", "amount", "recipientAccountId", "counterpartyId", "claimedFeeCounterparty","claimedFeeCustomer", "claimedFeeExpiresAt", "status", "counterpartyTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LoadMoneyMPGS.csv' DELIMITER ',';
COPY "public"."BCNUser" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUser.csv' DELIMITER ',';

COPY "public"."BCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNAccount.csv' DELIMITER ',';

COPY "public"."DefaultBCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/DefaultBCNAccount.csv' DELIMITER ',';

COPY "public"."Password" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Password.csv' DELIMITER ',';

COPY "public"."LeoLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LeoLLT.csv' DELIMITER ',';

COPY "public"."BCNUserLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUserLLT.csv' DELIMITER ',';

COPY "public"."Transaction"("id","amount","senderBcnAccountId","recipientBcnAccountId","resultedAt","exchangeRateMultiplier","exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Transaction.csv' DELIMITER ',' NULL 'null';


COPY "public"."SendMoneyWithinBCN"("id", "recipientBcnAccountId", "amount", "senderBcnAccountId","claimedFeeCustomer", "claimedFeeExpiresAt", "status","transactionId", "feeCustomerTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyWithinBCN.csv' DELIMITER ',';

COPY "public"."SendMoneyAccountToAccount"("id", "amount", "senderBcnAccountId", "recipientBcnAccountId", "claimedFeeCustomer","confirmationExpiresAt", "status", "transactionId","feeCustomerTransactionId", "exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyAccountToAccount.csv' DELIMITER ',';

COPY "public"."CounterpartyTransaction"("id", "counterpartyFinalTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/CounterPartyTransaction.csv' DELIMITER ',';

COPY "public"."LoadMoneyMPGS"("id", "amount", "recipientAccountId", "counterpartyId", "claimedFeeCounterparty","claimedFeeCustomer", "claimedFeeExpiresAt", "status", "counterpartyTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LoadMoneyMPGS.csv' DELIMITER ',';
COPY "public"."BCNUser" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUser.csv' DELIMITER ',';

COPY "public"."BCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNAccount.csv' DELIMITER ',';

COPY "public"."DefaultBCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/DefaultBCNAccount.csv' DELIMITER ',';

COPY "public"."Password" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Password.csv' DELIMITER ',';

COPY "public"."LeoLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LeoLLT.csv' DELIMITER ',';

COPY "public"."BCNUserLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUserLLT.csv' DELIMITER ',';

COPY "public"."Transaction"("id","amount","senderBcnAccountId","recipientBcnAccountId","resultedAt","exchangeRateMultiplier","exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Transaction.csv' DELIMITER ',' NULL 'null';


COPY "public"."SendMoneyWithinBCN"("id", "recipientBcnAccountId", "amount", "senderBcnAccountId","claimedFeeCustomer", "claimedFeeExpiresAt", "status","transactionId", "feeCustomerTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyWithinBCN.csv' DELIMITER ',';

COPY "public"."SendMoneyAccountToAccount"("id", "amount", "senderBcnAccountId", "recipientBcnAccountId", "claimedFeeCustomer","confirmationExpiresAt", "status", "transactionId","feeCustomerTransactionId", "exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyAccountToAccount.csv' DELIMITER ',';

COPY "public"."CounterpartyTransaction"("id", "counterpartyFinalTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/CounterPartyTransaction.csv' DELIMITER ',';

COPY "public"."LoadMoneyMPGS"("id", "amount", "recipientAccountId", "counterpartyId", "claimedFeeCounterparty","claimedFeeCustomer", "claimedFeeExpiresAt", "status", "counterpartyTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LoadMoneyMPGS.csv' DELIMITER ',';
COPY "public"."BCNUser" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUser.csv' DELIMITER ',';

COPY "public"."BCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNAccount.csv' DELIMITER ',';

COPY "public"."DefaultBCNAccount" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/DefaultBCNAccount.csv' DELIMITER ',';

COPY "public"."Password" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Password.csv' DELIMITER ',';

COPY "public"."LeoLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LeoLLT.csv' DELIMITER ',';

COPY "public"."BCNUserLLT" FROM '/var/lib/bcn/postgresql/db_data/csvFiles/BCNUserLLT.csv' DELIMITER ',';

COPY "public"."Transaction"("id","amount","senderBcnAccountId","recipientBcnAccountId","resultedAt","exchangeRateMultiplier","exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/Transaction.csv' DELIMITER ',' NULL 'null';


COPY "public"."SendMoneyWithinBCN"("id", "recipientBcnAccountId", "amount", "senderBcnAccountId","claimedFeeCustomer", "claimedFeeExpiresAt", "status","transactionId", "feeCustomerTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyWithinBCN.csv' DELIMITER ',';

COPY "public"."SendMoneyAccountToAccount"("id", "amount", "senderBcnAccountId", "recipientBcnAccountId", "claimedFeeCustomer","confirmationExpiresAt", "status", "transactionId","feeCustomerTransactionId", "exchangeRateId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/SendMoneyAccountToAccount.csv' DELIMITER ',';

COPY "public"."CounterpartyTransaction"("id", "counterpartyFinalTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/CounterPartyTransaction.csv' DELIMITER ',';

COPY "public"."LoadMoneyMPGS"("id", "amount", "recipientAccountId", "counterpartyId", "claimedFeeCounterparty","claimedFeeCustomer", "claimedFeeExpiresAt", "status", "counterpartyTransactionId") FROM '/var/lib/bcn/postgresql/db_data/csvFiles/LoadMoneyMPGS.csv' DELIMITER ',';
