BEGIN {
        RTT_new = 0
        RTT_old = 0
        flag = 0
        alpha = 0.25
}



{
        if(flag == 0){
                RTT_old = $2
                flag = 1
        }

        RTT_new = alpha*$2 + (1 - alpha)*RTT_old
        print $1, RTT_new
        RTT_old = RTT_new
}

END {
}