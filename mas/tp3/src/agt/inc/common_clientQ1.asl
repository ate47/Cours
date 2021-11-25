

+delay(LimitCompany)[source(Company)]:
    delay_limit(CurrentLimit) & CurrentLimit > LimitCompany
    <-
        // Update the clientQ1a.asl and clientQ1b.asl files so that the clients 
        // update their delay_limit belief based on the delay requested by the company
        // println(LimitCompany, " c=", Company)
        -+delay_limit(LimitCompany);
        .

