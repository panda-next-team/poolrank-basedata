<?php

namespace App\Admin\Models;

use Illuminate\Database\Eloquent\Model;
use App\Observers\ModelObserver;

class Algorithm extends Model
{
    protected $table = 'algorithm';


    public function powCoins()
    {
        return $this->hasMany(POWCoin::class, "id", "algorithm_id");
    }

    public static function boot()
    {
        parent::boot();
        $class = get_called_class();
        $class::observe(new ModelObserver());
    }
}
